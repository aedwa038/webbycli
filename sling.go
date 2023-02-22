package main

import (
	"github.com/dghubble/sling"
	"net/http"
)

type Fetcher struct {
	sling *sling.Sling
}

func newFetcher(baseURL string, httpClient *http.Client) *Fetcher {
	return &Fetcher{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

func (s Fetcher) Fetch(term string, params Params) ([]Entry, error) {
	entries := new([]Entry)
	_, err := s.sling.New().Path(term).QueryStruct(params).ReceiveSuccess(entries)
	return *entries, err
}
