# Five important questions about this code

- Go CLI

go run: compiles and execute one or two files or three or four a handful of files and instantly executes the result.

go build: compiles a bunch of go source code files

go fmt: formats all the code in each file in current directory

go install: compiles and "installs" a package

go get: downloads the raw source code of someone else's package

go test: runs any tests associated with the current project

- Package main

package == project == workspace

A package is a coolection of common source code files

Executable: Generates a file that we can run
Reasuable: Code used as 'helpers'. Good place to put reusable logic

The name of the package that you use that determines whether you are making an executable or depency type package

- import "fmt"

So saying import fmt specifically means give my package main access to all of the code in all the functionality that is contained inside of this other package called fmt.

- what is 'func'

func = function