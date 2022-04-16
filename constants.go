package main

import "fmt"

func main() {

	// const pi = 3.14 // untyped constants
	const pi float32 = 3.14 // typed constants
	// pi = 4.1 //cannot assign to pi (constant 3.14 of type float32)

	fmt.Println(pi)
}
