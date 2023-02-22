package main

type Entry struct {
	MetaData MetaDataInfo `json:"meta"`
	Fl       string       `json:"fl"`
	ShortDef []string     `json:"shortdef"`
	HeadWord HeadWordInfo `json:"hwi"`
}

type HeadWordInfo struct {
	Hw  string `json:"hw"`
	Prs []PRS  `json:"prs"`
}

type PRS struct {
	Mw    string    `json:"mw"`
	Sound SoundInfo `json:sound`
}

type SoundInfo struct {
	Audio string `json:audio`
}

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

type fetcher interface {
	Fetch(term string, params Params) ([]Entry, error)
}

type Params struct {
	Key string `url:"key,omitempty"`
}

// NewWebsterService returns a new .
func NewWebsterService(baseURL string, client fetcher, params Params) *WebsterService {
	return &WebsterService{
		fetcher: client,
		p:       params,
	}
}

// List returns the authenticated user's issues across repos and orgs.
func (s *WebsterService) ListDefs(term string) ([]Result, error) {
	entries, err := s.fetcher.Fetch(term, s.p)
	if err != nil {
		return []Result{}, err
	}
	finalEntries := filterEntries(entries, term)
	defs := mapToDefs(finalEntries)
	return defs, err
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func filterEntries(entries []Entry, term string) []Entry {
	e := make([]Entry, 0)
	for _, entry := range entries {
		if entry.MetaData.Id != "" && contains(entry.MetaData.Stems, term) {
			e = append(e, entry)
		}

	}
	return e
}

type Result struct {
	//metaData MetaDataInfo
	Label string
	Defs  []string
	Prs   string
}

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
