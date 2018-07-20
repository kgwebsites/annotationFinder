# Annotation Finder

Executable file which finds all your // TODO and // FIXME annotation comments within your directory and appends them to your README.md file.

## Installation

You have 3 options here.

1. EASIEST - Download the `annotations` (mac / linux) or `annotations.exe` (windows) binary file and execute it in a linux operating system.
2. You can clone the package somewhere and access the `annotations` binary file within the package.
3. You can install with NPM `npm i --dev annotationfinder` or Yarn `yarn add --dev annotationfinder` and you will have access to the entire repo, in which you will want to use the appropriate binary file.

## Usage

To use simply call the executable in this package `annotations` (mac / linux) or `annotations.exe` (windows).

You can change the directory set to search for annotations by using the directory flag:
mac / linux
```$ ./annotations -d="./testdata" OR $ ./annotations -directory="./testdata"```
windows
```./annotations.exe -d="./testdata" OR $ `./annotations.exe -directory="./testdata"```

You can also use this for a single file by using the file flag:
mac / linux
```$ ./annotations -f="./testdata/foo"` OR $ ./annotations -file="./testdata/foo"```
windows
```$ ./annotations.exe -f="./testdata/foo"` OR $ ./annotations.exe -file="./testdata/foo"```

If you want your annotations to be appended to the end of the README.md file in the current directory in markdown format, pass in the append flag:
mac / linux
```$ ./annotations -a OR ./annotations -append```
windows
```$ ./annotations.exe -a OR $ ./annotations.exe -append```

To specify the file which the appending is done, pass in the output flag:
mac / linux
```$ ./annotations -a -o "SECONDARYREADME.md" OR $ ./annotations -append -output "SECONDARYREADME.md"```
windows
```$ ./annotations.exe -a -o "SECONDARYREADME.md" OR $ ./annotations.exe -append -output "SECONDARYREADME.md"```

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