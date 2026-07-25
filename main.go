package main

import (
	"fmt"
	"log"
	"os"
	"tag_search/search"
	"tag_search/ui"
)

func main() {
	config := ui.ParseFlags()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	targets, err := search.CollectTags(dir)
	if err != nil {
		log.Fatal(err)
	}

	var filtered []search.TagSet
	for _, tagSet := range targets {
		if matchTags(tagSet, config) {
			filtered = append(filtered, tagSet)
		}
	}
	for _, ts := range filtered {
		fmt.Println(ts.Path)
	}
}
