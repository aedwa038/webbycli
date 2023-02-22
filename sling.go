package main

import (
	"github.com/dghubble/sling"
	"net/http"
)

// Fetcher object used for getting word defnitions
type Fetcher struct {
	sling *sling.Sling
}

func newFetcher(baseURL string, httpClient *http.Client) *Fetcher {
	return &Fetcher{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

// Fetch gets a list of definitions for a search term.
func (s Fetcher) Fetch(term string, params Params) ([]Entry, error) {
	entries := new([]Entry)
	_, err := s.sling.New().Path(term).QueryStruct(params).ReceiveSuccess(entries)
	return *entries, err
}
