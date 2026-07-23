package main

import (
	"bufio"
	"fmt"
	"io/fs"
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
		} //for permision denied and stuf like that
		if d.IsDir() || d.Name() != ".tags.txt" {
			return nil // skip, keep walking
		}

		tagSet, err := readLinesToStruct(path) // path is already full/correct
		if err != nil {
			return err
		}

		results = append(results, tagSet)
		return nil
	})

	return results, err
}
func main() {
	set, err := readLinesToStruct("blocklist.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(set)
}
