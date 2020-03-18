package main

import (
	"fmt"
	"regexp"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// DirStructure is used to parse changes commit
const DirStructure = `(?P<fullpath>(?P<dirpath>providers\/(?P<providers>[\w]*).*\/)(?P<filename>[\w|\W]*.tf))`

func main() {
	var matcher *regexp.Regexp
	//var finalDir string
	matcher = regexp.MustCompile(DirStructure)

	repo, _ := git.PlainOpen(".")
	ref, _ := repo.Head()
	commit, _ := repo.CommitObject(ref.Hash())
	fileStats := object.FileStats{}

	fileStats, _ = commit.Stats()

	filePaths := []string{}

	for _, fileStat := range fileStats {
		filePaths = append(filePaths, fileStat.Name)
	}

	f := make(map[string]string)
	pr := []string{}
	for _, filePath := range filePaths {
		if !strings.Contains(filePath, "/") {
			continue
		}

		// check if the changes is in the same directory
		matches := matcher.FindStringSubmatch(filePath)

		if len(matches) == 0 {
			continue
		} else {
			for i, name := range matcher.SubexpNames() {
				if name == "" {
					continue
				}
				f[name] = matches[i]
			}
		}

		pr = append(pr, f["dirpath"])
	}
	dir := removeDupes(pr)
	fmt.Println(dir)
}

func removeDupes(folder []string) []string {
	e := map[string]bool{}

	for i := range folder {
		e[folder[i]] = true
	}

	result := []string{}
	for key, _ := range e {
		result = append(result, key)
	}

	return result
}
