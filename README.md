# Go

Go is a statically typed, compiled programming language designed at Google by **Robert Griesemer**, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

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
	fmt.Println("Hello World!")
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
