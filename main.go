package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var term = flag.String("term", "", "Word to be looked up")
var configFile = flag.String("config", "config.yaml", "Word to be looked up")

const baseURL = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/"

func main() {
	flag.Parse()
	if *term == "" {
		log.Fatalf("No search term provided.")
	}
	config, err := getConfig(*configFile)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	if config.Key == "" {
		log.Fatalf("Unable to find api key")
	}
	params := Params{Key: config.Key}

	client := NewWebsterService(baseURL, newFetcher(baseURL, &http.Client{}), params)
	entries, err := client.ListDefs(*term)
	if err != nil {
		log.Fatalf("Unable to find term. %v", err)
	}

	if len(entries) < 1 {
		fmt.Println("No dictionary entries found")
	}
	fmt.Println(*term + "\n")
	output(entries)

}

func output(defs []Result) {
	for _, d := range defs {
		var sb strings.Builder
		fmt.Fprintf(&sb, "`%s`", d.Prs)
		fmt.Fprintf(&sb, "(%s):\n", d.Label)

		for _, i := range d.Defs {
			fmt.Fprintf(&sb, "\t%s\n", i)
		}
		fmt.Println(sb.String())
	}
}
