# Go

- Go is a statically typed, compiled programming language designed at Google by **Robert Griesemer**, Rob Pike, and Ken Thompson.
- IT is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.
- It is an open source programming language that makes it easy to build simple, reliable, and efficient software.

## Characteristics of go
- `Simple and readable` syntax like dynamically typed programming language like python
- `Efficency and safety` of a lower-level language, staically typed language like C++
- Fast build time, start up and run
- Requires fewer resources like cpu and ram to run
- Compiler into simple binary(machine code)
- Faster than interpreter language like python
- Consistent across diffrent OS

## Commands in go

```
steffytw@steffytw-HP-240-G7-Notebook-PC:~ $go
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	work        workspace maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	packages        package lists and patterns
	private         configuration for downloading non-public code
	testflag        testing flags
	testfunc        testing functions
	vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.

```

## Basic structure

A Go program basically consists of the following parts −

- Package Declaration
- Import Packages
- Functions
- Variables
- Statements and Expressions
- Comments

```
package main

import "fmt"

func main() {
	fmt.Println("Hello World!") //Hello World!
}
```

## main and init function in Golang

The Go language reserve two functions for special purpose and the functions are main() and init() function.

### main() function
- In Go language, the main package is a special package which is used with the programs that are executable and this package contains main() function. 
- The main() function is a special type of function and it is the entry point of the executable programs. 
- It does not take any argument nor return anything. 
- Go automatically call main() function, so there is no need to call main() function explicitly and every executable program must contain single main package and main() function.

### init() Function

- init() function is just like the main function, does not take any argument nor return anything. 
- This function is present in every package and this function is called when the package is initialized. 
- This function is declared implicitly, so you cannot reference it from anywhere and you are allowed to create multiple init() function in the same program and they execute in the order they are created. 
- You are allowed to create init() function anywhere in the program and they are called in lexical file name order (Alphabetical Order). 
- And allowed to put statements if the init() function, but always remember to init() function is executed before the main() function call, so it does not depend to main() function. 
- The main purpose of the init() function is to initialize the global variables that cannot be initialized in the global context.

## Package

Package can be of two types.

- *Executable package* – Only main is the executable package in GoLang. A .go file might belong to the main package present within a specific directory. We will see later how the directory name or the .go file name matters.  The main package will contain a main function that denotes the start of a program. On installing the main package it will create an executable in the $GOBIN directory.

- *Utility package*– Any package other than the main package is a utility package. It is not self-executable. It just contains the utility function and other utility things that can be utilized by an executable package.
