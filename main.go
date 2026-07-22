package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func readLinesToSet(path string) (map[string]struct{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	set := make(map[string]struct{})
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		set[strings.ToLower(line)] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return set, nil
}

type TagSet struct {
	Path string
	Tags map[string]struct{}
}

func collectTagSets(root string) ([]TagSet, error) {
	var results []TagSet

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		} //for permision denied and stuf like that
		if d.IsDir() || d.Name() != ".tags.txt" {
			return nil // skip, keep walking
		}

		tags, err := readLinesToSet(path) // path is already full/correct
		if err != nil {
			return err
		}

		results = append(results, TagSet{Path: path, Tags: tags})
		return nil
	})

	return results, err
}
func main() {
	set, err := readLinesToSet("blocklist.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(len(set), "lines")
}
