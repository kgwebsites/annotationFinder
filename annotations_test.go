package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	main()
	os.Exit(m.Run())
}

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
		t.Errorf("Expected to find '/testdata/foo' within the files")
	}

	// Test for file -flag
	fileflag = "./testdata/foo"
	fi := findFiles("./testdata")
	if fi[0] != "./testdata/foo" {
		t.Errorf("Expected to return flagged file")
	}
	fileflag = ""
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

func read() (file []byte, err error) {
	return ioutil.ReadFile("README.md")
}

func TestOutput(t *testing.T) {
	old, _ := read()
	output("foo")
	new, _ := read()
	end := new[len(new)-3:]

	if string(end) != "foo" {
		t.Errorf("Expected the README.md file to have been appended, 'foo'")
	} else {
		os.Remove("README.md")
		// Create new README.md file
		ioutil.WriteFile("README.md", []byte(old), 0666)
	}

}
