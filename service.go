package main

// Entry represents a definition returned by the api
type Entry struct {
	MetaData MetaDataInfo `json:"meta"`
	Fl       string       `json:"fl"`
	ShortDef []string     `json:"shortdef"`
	HeadWord HeadWordInfo `json:"hwi"`
}

// HeadWordInfo represents the word be defined includes pronounciation.
type HeadWordInfo struct {
	Hw  string `json:"hw"`
	Prs []PRS  `json:"prs"`
}

// PRS represents the pronounciation of a word being defined
type PRS struct {
	Mw    string    `json:"mw"`
	Sound SoundInfo `json:sound`
}

// SoundInfo represents the audio file for the word being defined
type SoundInfo struct {
	Audio string `json:audio`
}

// MetadataInfo represents metadata on the defined word.
type MetaDataInfo struct {
	Id    string   `json:"id"`
	UUID  string   `json:"uuid"`
	Stems []string `json:"stems"`
}

// WebsterService provides methods for reading Webster dictionay definitions.
type WebsterService struct {
	fetcher fetcher
	p       Params
}

// fetcher is a genric interface for communitcating with the webster api
type fetcher interface {
	Fetch(term string, params Params) ([]Entry, error)
}

// Params represent the parameters needed to communicate with the api
type Params struct {
	Key string `url:"key,omitempty"`
}

// NewWebsterService returns a new WebsterService for fetching definitions
func NewWebsterService(baseURL string, client fetcher, params Params) *WebsterService {
	return &WebsterService{
		fetcher: client,
		p:       params,
	}
}

// ListDefs returns a list of defnitions found for a word or search term
func (s *WebsterService) ListDefs(term string) ([]Result, error) {
	entries, err := s.fetcher.Fetch(term, s.p)
	if err != nil {
		return []Result{}, err
	}
	finalEntries := filterEntries(entries, term)
	defs := mapToDefs(finalEntries)
	return defs, err
}

// checks if a word is contained within an array.
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// filterEntries filters entries that are invalid or not neccsary
func filterEntries(entries []Entry, term string) []Entry {
	e := make([]Entry, 0)
	for _, entry := range entries {
		if entry.MetaData.Id != "" && contains(entry.MetaData.Stems, term) {
			e = append(e, entry)
		}

	}
	return e
}

// Result represents a definition of a word
type Result struct {
	//metaData MetaDataInfo
	Label string
	Defs  []string
	Prs   string
}

//mapToDefs maps Entries to results being returned
func mapToDefs(entries []Entry) []Result {
	var defs []Result
	for _, entry := range entries {
		def := Result{}
		if len(entry.HeadWord.Prs) != 0 {
			p := entry.HeadWord.Prs[0]
			if p.Mw != "" {
				def.Prs = p.Mw
			}
		} else if entry.HeadWord.Hw != "" {
			def.Prs = entry.HeadWord.Hw
		}
		if entry.Fl != "" {
			def.Label = entry.Fl
		}

		for _, d := range entry.ShortDef {
			def.Defs = append(def.Defs, d)
		}
		defs = append(defs, def)
	}

	return defs
}
