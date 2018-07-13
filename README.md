# Annotation Finder

Executable file which finds all your // TODO , // FIXME , // REFACTOR annotation comments within your directory and appends them to your README.md file.

## Installation

You have 3 options here.

1. EASIEST - Download the `annotations` (mac / linux) or `annotations.exe` (windows) binary file and execute it in a linux operating system.
2. You can clone the package somewhere and access the `annotations` binary file within the package.
3. You can install with NPM `npm i --dev annotationfinder` or Yarn `yarn add --dev annotationfinder` and you will have access to the entire repo, in which you will want to use the appropriate binary file.

## Usage

To use simply call the executable in this package `annotations` (mac / linux) or `annotations.exe` (windows).

You can change the directory set to search for annotations by using the dir flag. For example:
`./annotations -dir="./testdata"` (mac / linux)
`./annotations.exe -dir="./testdata"` (windows)

You can also use this for a single file by using the file flag.
`./annotations -file="./testdata/foo"` (mac / linux)
`./annotations -file="./testdata/foo"` (windows)

## Contributing

All contributions are welcome and appreciated!
The easiest way is for you to fork your own version, make a new feature branch (example: feature/coverage-increase-50-75) and then make a PR.

All code lies in annotations.go, make your changes in there, and build out both the linux and windows binaries like so:
On a mac / linux:
```
go build -o annotations
env GOOS=windows GOARCH=amd64 go build -o annotations.exe
```
On a windows:
```
go build -o annotations.exe
env GOOS=linux GOARCH=arm go build -o annotations
```
The -o flag is important to make sure the executable is named properly and not the default package name.