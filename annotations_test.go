package main

import (
	"strings"
	"testing"
)

func TestFindFiles(t *testing.T) {
	f := findFiles("./testdata")

	if f[0] != "./testdata" {
		t.Errorf("Expected the first file to be the directory, './testdata', got %v", f[0])
	}

	e := false
	for _, i := range f {
		if i == "testdata/file1" {
			e = true
		}
	}

	if e != true {
		t.Errorf("Expected to find '/testdata/file1' within the files")
	}
}

func TestSkipped(t *testing.T) {
	if skipped("node_modules", true) == false {
		t.Errorf("Expected to skip node_modules folder")
	}
	if skipped("annotations", false) == false {
		t.Errorf("Expected to skip annotations file")
	}
}

func TestFindAnnotations(t *testing.T) {
	td := findAnnotations(findFiles("./testdata"))["TODOS"][0].todo
	f := findAnnotations(findFiles("./testdata"))["FIXME"][0].todo
	r := findAnnotations(findFiles("./testdata"))["REFACTOR"][0].todo

	if td != "foo" {
		t.Errorf("Expected to find 'foo' in the TODOS")
	}

	if f != "bar" {
		t.Errorf("Expected to find 'bar' in the FIXME")
	}

	if r != "baz" {
		t.Errorf("Expected to find 'baz' in the REFACTOR")
	}
}

func TestBuildList(t *testing.T) {
	f := findFiles("./testdata")
	a := findAnnotations(f)
	l := buildList(a)
	if strings.Contains(l, "[foo]") != true {
		t.Errorf("Expected the list to have foo listed in the TODOS")
	}
}
