package main

import "fmt"

func main() {
	fmt.Println("Go is statically typed")
}

func init() { //init() function is executed before the main() function call
	fmt.Println("Welcome to Go!")
}
