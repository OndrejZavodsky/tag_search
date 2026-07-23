package main

import "sort"

func matchTags(requested []string, tagSet TagSet, cfg Config) bool {
	tags := dedupe(tagSet.Tags)
	matched := 0
	for _, tag := range tags {
		for _, req := range requested {
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
