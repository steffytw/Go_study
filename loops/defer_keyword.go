package main

import "fmt"

func main() {
	defer test()
	defer func() { fmt.Println("In inline defer") }()
	fmt.Println("Start main")
	f1()
	fmt.Println("Finish main")
}
func test() {
	fmt.Println("In Defer")
}

func f1() {
	defer fmt.Println("Defer in f1")
	fmt.Println("Start f1")
	f2()
	fmt.Println("Finish f1")
}

func f2() {
	defer fmt.Println("Defer in f2")
	fmt.Println("Start f2")
	fmt.Println("Finish f2")
}
