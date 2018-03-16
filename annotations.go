package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dirflag string
var fileflag string

func init() {
	flag.StringVar(&dirflag, "dir", "./", "Directory of files you wish to search for annotations")
	flag.StringVar(&fileflag, "file", "", "Single file you wish to search for annotations")
}

type annotation struct {
	path string
	todo string
}

func main() {
	flag.Parse()
	f := findFiles(dirflag)
	a := findAnnotations(f)
	l := buildList(a)
	appendReadme(l)
}

func findFiles(dir string) []string {
	files := []string{}

	if fileflag != "" {
		files = append(files, fileflag)
	} else {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("%v\n", err)
				return err
			}
			if !skipped(info.Name(), true) {
				if !skipped(info.Name(), false) {
					files = append(files, path)
				}
			} else {
				return filepath.SkipDir
			}
			return nil
		})

		if err != nil {
			fmt.Printf("error walking the path %q: %v\n", dir, err)
		}
	}

	return files
}

func skipped(name string, dir bool) bool {
	skipDir := []string{"node_modules", ".git", "coverage"}
	skipFiles := []string{"annotations.go", "annotations_test.go", "annotations", "README.md", "Jenkinsfile", "cover.out", "debug", "debug.test", "LICENSE"}
	ref := skipFiles

	if dir {
		ref = skipDir
	}

	for _, s := range ref {
		if s == name {
			return true
		}
	}
	return false
}

func findAnnotations(files []string) map[string][]annotation {
	annotations := map[string][]annotation{
		"TODOS":    {},
		"FIXME":    {},
		"REFACTOR": {},
	}
	for _, file := range files {
		openFile, _ := os.Open(file)
		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			if strings.Contains(string(scanner.Text()), "TODO:") {
				annotations["TODOS"] = append(annotations["TODOS"],
					annotation{path: file, todo: strings.TrimPrefix(string(scanner.Text()), "// TODO: ")})
			}
			if strings.Contains(string(scanner.Text()), "FIXME:") {
				annotations["FIXME"] = append(annotations["FIXME"],
					annotation{path: file, todo: strings.TrimPrefix(string(scanner.Text()), "// FIXME: ")})
			}
			if strings.Contains(string(scanner.Text()), "REFACTOR") {
				annotations["REFACTOR"] = append(annotations["REFACTOR"],
					annotation{path: file, todo: strings.TrimPrefix(string(scanner.Text()), "// REFACTOR: ")})
			}
		}
	}
	return annotations
}

func buildList(notes map[string][]annotation) string {
	list := "\n## ANNOTATIONS\n"
	for i, n := range notes {
		if len(n) > 0 {
			list = list + "### " + i + ":\n"
			for _, a := range n {
				list = list + "* [" + a.todo + "](" + a.path + ")\n"
			}
			list = list + "\n"
		}
	}
	return list
}

func appendReadme(list string) {
	// Create README.md if nonexistant
	f, err := os.OpenFile("README.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		// Save old file data
		old, e := ioutil.ReadFile("README.md")
		if e != nil {
			fmt.Println(e)
		}
		// Recreate README.md string with old list removed
		removed := strings.Split(string(old), "\n## ANNOTATIONS")
		// Append new list to new README.md String.
		readme := removed[0] + list
		// Remove current README.md file
		os.Remove("README.md")
		// Create new README.md file
		ioutil.WriteFile("README.md", []byte(readme), 0666)
	} else {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
