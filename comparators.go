package main

import (
	"sort"
	"tag_search/search"
	"tag_search/ui"
)

func matchTags(tagSet search.TagSet, cfg ui.Config) bool {
	tags := dedupe(tagSet.Tags)
	matched := 0
	for _, tag := range tags {
		for _, req := range cfg.Tags {
			if tag == req {
				matched++
			}
		}
	}
	return matched >= cfg.MatchCount
}

func dedupe(s []string) []string {
	sort.Strings(s)
	out := s[:0]
	for i, v := range s {
		if i == 0 || v != s[i-1] {
			out = append(out, v)
		}
	}
	return out
}
