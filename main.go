package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func readLinesToStruct(path string) (TagSet, error) {
	f, err := os.Open(path)
	if err != nil {
		return TagSet{}, err
	}
	defer f.Close()

	var tags []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tags = append(tags, line)
	}
	if err := scanner.Err(); err != nil {
		return TagSet{}, err
	}
	tagSet := TagSet{Path: path, Tags: tags}
	return tagSet, nil
}

type TagSet struct {
	Path string
	Tags []string
}

func collectTags(root string) ([]TagSet, error) {
	var results []TagSet

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || d.Name() != ".tags.txt" {
			return nil
		}

		tagSet, err := readLinesToStruct(path)
		if err != nil {
			return err
		}

		results = append(results, tagSet)
		return nil
	})

	return results, err
}
func main() {
	config := ParseFlags()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	targets, err := collectTags(dir)
	if err != nil {
		log.Fatal(err)
	}

	var filtered []TagSet
	for _, tagSet := range targets {
		if matchTags(tagSet, config) {
			filtered = append(filtered, tagSet)
		}
	}
	for _, ts := range filtered {
		fmt.Println(ts.Path)
	}
}
