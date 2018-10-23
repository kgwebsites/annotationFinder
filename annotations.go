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
var appendflag bool
var outputflag string
var rejectfixmeflag bool

func init() {
	flag.StringVar(&dirflag, "d", "./", "[shorthand] Directory of files you wish to search for annotations")
	flag.StringVar(&dirflag, "directory", "./", "[verbose] Directory of files you wish to search for annotations")
	flag.StringVar(&fileflag, "f", "", "[shorthand] Single file you wish to search for annotations")
	flag.StringVar(&fileflag, "file", "", "[verbose] Single file you wish to search for annotations")
	flag.BoolVar(&appendflag, "a", false, "[shorthand] Pass in this flag if you want to append annotations to a markdown file")
	flag.BoolVar(&appendflag, "append", false, "[verbose] Pass in this flag if you want to append annotations to a markdown file")
	flag.StringVar(&outputflag, "o", "README.md", "[shorthand] Markdown file you wish append annotations to")
	flag.StringVar(&outputflag, "output", "README.md", "[verbose] Markdown file you wish append annotations to")
	flag.BoolVar(&rejectfixmeflag, "rf", false, "[shorthand] Option to return error if a fix me is found.")
	flag.BoolVar(&rejectfixmeflag, "rejectfixme", false, "[verbose] Option to return error if a fix me is found.")
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
	if appendflag {
		output(l)
	}
	if rejectfixmeflag {
		rejectFixme(a)
	}
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
	green := color.New(color.FgHiGreen).SprintFunc()
	red := color.New(color.FgHiRed).SprintFunc()
	blue := color.New(color.FgHiCyan).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	for i, n := range notes {
		if len(n) > 0 {
			fmt.Printf("%v:\n", i)
			list = list + "### " + i + ":\n"
			for _, a := range n {
				list = list + "* [" + a.item + "](" + a.path + ")\n"
				if i == "TODO" {
					fmt.Printf("%s %s %s %s\n", white("*"), green(a.item), white("-"), blue(a.path))
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

func output(list string) {
	// Create README.md if nonexistant
	f, err := os.OpenFile(outputflag, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		// Save old file data
		old, e := ioutil.ReadFile(outputflag)
		if e != nil {
			fmt.Println(e)
		}
		// Recreate README.md string with old list removed
		removed := strings.Split(string(old), "\n## ANNOTATIONS")
		// Append new list to new README.md String.
		readme := removed[0] + list
		// Remove current README.md file
		os.Remove(outputflag)
		// Create new README.md file
		ioutil.WriteFile(outputflag, []byte(readme), 0666)
	} else {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func rejectFixme(notes map[string][]annotation) {
	fixmeExist := false
	red := color.New(color.FgHiRed).SprintFunc()
	blue := color.New(color.FgHiCyan).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	list := "Fix the following FIXMEs before continuing: \n\n"
	for i, n := range notes {
		for _, a := range n {
			if i == "FIXME" {
				fixmeExist = true
				list = fmt.Sprintf("%s %s %s %s %s\n", list, white("*"), red(a.item), white("-"), blue(a.path))
			}

		}
	}
	if fixmeExist {
		log.Fatalf("\n\n========================================================================\n\n%s \n========================================================================\n\n", list)
	}
}
