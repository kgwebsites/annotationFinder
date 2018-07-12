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

	"github.com/fatih/color"
)

var dirflag string
var fileflag string

func init() {
	flag.StringVar(&dirflag, "dir", "./", "Directory of files you wish to search for annotations")
	flag.StringVar(&fileflag, "file", "", "Single file you wish to search for annotations")
}

type annotation struct {
	path string
	item string
}

func main() {
	flag.Parse()
	f := findFiles(dirflag)
	a := findAnnotations(f)
	l := buildAndLogList(a)
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
		"TODO":  {},
		"FIXME": {},
	}
	for _, file := range files {
		openFile, _ := os.Open(file)
		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			line := string(scanner.Text())
			if strings.Contains(line, "TODO:") {
				todo := strings.Split(line, "TODO: ")
				annotations["TODO"] = append(annotations["TODO"],
					annotation{path: file, item: todo[1]})
			}
			if strings.Contains(line, "FIXME:") {
				fixme := strings.Split(line, "FIXME: ")
				annotations["FIXME"] = append(annotations["FIXME"],
					annotation{path: file, item: fixme[1]})
			}
		}
	}
	return annotations
}

func buildAndLogList(notes map[string][]annotation) string {
	list := "\n## ANNOTATIONS\n"
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgCyan).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	for i, n := range notes {
		if len(n) > 0 {
			fmt.Printf("%v:\n", i)
			list = list + "### " + i + ":\n"
			for _, a := range n {
				list = list + "* [" + a.item + "](" + a.path + ")\n"
				if i == "TODO" {
					fmt.Printf("%s %s %s %s\n", white("*"), green(a.item), white(" - "), blue(a.path))
				}
				if i == "FIXME" {
					fmt.Printf("%s %s %s %s\n", white("*"), red(a.item), white("-"), blue(a.path))
				}
			}
			println()
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
