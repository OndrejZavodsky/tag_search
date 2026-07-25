package ui

import (
	"flag"
	"strings"
)

type Config struct {
	Tags       []string
	MatchCount int
}

func ParseFlags() Config {
	tagsFlag := flag.String("tags", "", "comma-separated tags, e.g. tag1,tag2")
	fullMatch := flag.Bool("f", false, "require full match")
	matchCount := flag.Int("m", 1, "number of matching entries")

	flag.Parse()

	var tags []string
	if *tagsFlag != "" {
		tags = strings.Split(*tagsFlag, ",")
	}
	if *fullMatch {
		return Config{Tags: tags,
			MatchCount: len(tags)}
	}
	return Config{
		Tags:       tags,
		MatchCount: *matchCount,
	}
}
