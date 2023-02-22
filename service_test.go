package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	var tests = []struct {
		words []string
		term  string
		want  bool
		desc  string
	}{
		{
			desc:  "Matching test case",
			words: []string{"grotesque", "grotesques"},
			term:  "grotesque",
			want:  true,
		},
		{
			desc:  "Matching test case 2",
			words: []string{"grotesque", "grotesquely", "grotesqueness", "grotesquenesses", "grotesquer", "grotesquest"},
			term:  "grotesque",
			want:  true,
		},
		{
			desc:  "No Matches test case",
			words: []string{"coach dog", "coach dogs", "hackney coach", "hackney coaches"},
			term:  "coach",
			want:  false,
		},
		{
			desc:  "Emptystring test case",
			words: []string{""},
			term:  "coach",
			want:  false,
		},
		{
			desc:  "Empty string test case",
			words: []string{"coaches"},
			term:  "",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ans := contains(test.words, test.term)
			if ans != test.want {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})

	}

}
func TestFilter(t *testing.T) {
	var tests = []struct {
		entries []Entry
		want    []Entry
		term    string
		desc    string
	}{
		{
			desc: "Normal test case",
			entries: []Entry{
				Entry{
					MetaData: MetaDataInfo{
						Id:    "word1",
						Stems: []string{"wording"},
					},
				},
			},
			want: []Entry{
				Entry{
					MetaData: MetaDataInfo{
						Id:    "word1",
						Stems: []string{"wording"},
					},
				},
			},
			term: "wording",
		},
		{
			desc: "Complicated test case",
			entries: []Entry{
				Entry{
					MetaData: MetaDataInfo{
						Id:    "word:1",
						Stems: []string{"word", "a good word", "good word", "in a word", "in so many words", "of few words"},
					},
				},
				Entry{
					MetaData: MetaDataInfo{
						Id:    "word:2",
						Stems: []string{"word", "worded", "wording", "words"},
					},
				},
				Entry{
					MetaData: MetaDataInfo{
						Id:    "word class",
						Stems: []string{"word class", "word classes"},
					},
				},
			},
			want: []Entry{
				Entry{
					MetaData: MetaDataInfo{
						Id:    "word:1",
						Stems: []string{"word", "a good word", "good word", "in a word", "in so many words", "of few words"},
					},
				},
				Entry{
					MetaData: MetaDataInfo{
						Id:    "word:2",
						Stems: []string{"word", "worded", "wording", "words"},
					},
				},
			},
			term: "word",
		},

		{
			desc: "Empty ID test case",
			entries: []Entry{
				Entry{
					MetaData: MetaDataInfo{
						Id:    "",
						Stems: []string{"wording"},
					},
				},
			},
			want: []Entry{},
			term: "wording",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			out := filterEntries(test.entries, test.term)
			if !reflect.DeepEqual(out, test.want) {
				t.Errorf("got %v, want %v", out, test.want)
			}
		})

	}
}

func TestMapToDefs(t *testing.T) {
	var tests = []struct {
		entries []Entry
		want    []Result
		desc    string
	}{
		{
			entries: []Entry{
				Entry{
					MetaData: MetaDataInfo{
						Id:    "grotesque:1",
						Stems: []string{"grotesque", "grotesques"},
					},
					Fl: "noun",
					ShortDef: []string{
						"a style of decorative art characterized by fanciful or fantastic human and animal forms often interwoven with foliage or similar figures that may distort the natural into absurdity, ugliness, or caricature",
						"a piece of work in this style",
						"one that is grotesque",
					},
					HeadWord: HeadWordInfo{
						Hw: "gro*tesque",
						Prs: []PRS{
							PRS{
								Mw: "grō-ˈtesk",
							},
						},
					},
				},
				Entry{
					MetaData: MetaDataInfo{
						Id: "grotesque:2",
						Stems: []string{
							"grotesque",
							"grotesquely",
							"grotesqueness",
							"grotesquenesses",
							"grotesquer",
							"grotesquest",
						},
					},
					Fl: "adjective",
					ShortDef: []string{
						"of, relating to, or having the characteristics of the grotesque: such as",
						"fanciful, bizarre",
						"absurdly incongruous",
					},
					HeadWord: HeadWordInfo{
						Hw: "gro*tesque",
					},
				},
			},
			want: []Result{
				Result{
					Label: "noun",
					Defs: []string{
						"a style of decorative art characterized by fanciful or fantastic human and animal forms often interwoven with foliage or similar figures that may distort the natural into absurdity, ugliness, or caricature",
						"a piece of work in this style",
						"one that is grotesque",
					},
					Prs: "grō-ˈtesk",
				},
				Result{
					Label: "adjective",
					Defs: []string{
						"of, relating to, or having the characteristics of the grotesque: such as",
						"fanciful, bizarre",
						"absurdly incongruous",
					},
					Prs: "gro*tesque",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			out := mapToDefs(test.entries)
			if diff := cmp.Diff(test.want, out); diff != "" {
				t.Errorf("mapToDefs() mismatch (-want +got):\n%s", diff)
			}
		})

	}

}

func TestListDefs(t *testing.T) {
	var tests = []struct {
		term    string
		want    []Result
		wanterr bool
		fetcher fakeFetcher
		desc    string
	}{
		{
			desc:    "Normal test case",
			term:    "grotesque",
			wanterr: false,
			want: []Result{
				Result{
					Label: "noun",
					Defs: []string{
						"a style of decorative art characterized by fanciful or fantastic human and animal forms often interwoven with foliage or similar figures that may distort the natural into absurdity, ugliness, or caricature",
						"a piece of work in this style",
						"one that is grotesque",
					},
					Prs: "grō-ˈtesk",
				},
				Result{
					Label: "adjective",
					Defs: []string{
						"of, relating to, or having the characteristics of the grotesque: such as",
						"fanciful, bizarre",
						"absurdly incongruous",
					},
					Prs: "gro*tesque",
				},
			},
			fetcher: newfakeFetcher(false),
		},

		{
			desc:    "Failing test case",
			term:    "grotesque",
			wanterr: true,
			want: []Result{
				Result{
					Label: "noun",
					Defs: []string{
						"a style of decorative art characterized by fanciful or fantastic human and animal forms often interwoven with foliage or similar figures that may distort the natural into absurdity, ugliness, or caricature",
						"a piece of work in this style",
						"one that is grotesque",
					},
					Prs: "grō-ˈtesk",
				},
				Result{
					Label: "adjective",
					Defs: []string{
						"of, relating to, or having the characteristics of the grotesque: such as",
						"fanciful, bizarre",
						"absurdly incongruous",
					},
					Prs: "gro*tesque",
				},
			},
			fetcher: newfakeFetcher(true),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			client := NewWebsterService(baseURL, test.fetcher, Params{})
			out, err := client.ListDefs(test.term)
			if test.wanterr {
				if err == nil {
					t.Errorf("ListDefs() mismatch wanted error got %v", err)
				}
			} else if diff := cmp.Diff(test.want, out); diff != "" {
				t.Errorf("ListDefs() mismatch (-want +got):\n%s", diff)
			}
		})

	}

}

type fakeFetcher struct {
	entries []Entry
	fail    bool
}

func newfakeFetcher(f bool) fakeFetcher {
	return fakeFetcher{
		fail: f,
		entries: []Entry{
			Entry{
				MetaData: MetaDataInfo{
					Id:    "grotesque:1",
					Stems: []string{"grotesque", "grotesques"},
				},
				Fl: "noun",
				ShortDef: []string{
					"a style of decorative art characterized by fanciful or fantastic human and animal forms often interwoven with foliage or similar figures that may distort the natural into absurdity, ugliness, or caricature",
					"a piece of work in this style",
					"one that is grotesque",
				},
				HeadWord: HeadWordInfo{
					Hw: "gro*tesque",
					Prs: []PRS{
						PRS{
							Mw: "grō-ˈtesk",
						},
					},
				},
			},
			Entry{
				MetaData: MetaDataInfo{
					Id: "grotesque:2",
					Stems: []string{
						"grotesque",
						"grotesquely",
						"grotesqueness",
						"grotesquenesses",
						"grotesquer",
						"grotesquest",
					},
				},
				Fl: "adjective",
				ShortDef: []string{
					"of, relating to, or having the characteristics of the grotesque: such as",
					"fanciful, bizarre",
					"absurdly incongruous",
				},
				HeadWord: HeadWordInfo{
					Hw: "gro*tesque",
				},
			},
		},
	}
}

func (s fakeFetcher) Fetch(term string, params Params) ([]Entry, error) {
	if s.fail {
		return []Entry{}, fmt.Errorf("error reading response")
	}
	return s.entries, nil
}
