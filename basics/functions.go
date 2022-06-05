package main

import (
	"fmt"
)

func sampleMessage() {
	fmt.Println("Welcome to GO!")
}

func name(fname, mname, lname string) {
	fmt.Println("Hello!,", fname, mname, lname)

}
func contact(address string, phonenumber int) {
	fmt.Printf("Addres : %v \nPhone number: %v\n", address, phonenumber)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) (result int) {
	result = x - y
	return
}

func main() {
	// call the function
	sampleMessage()
	name("Lena", "John", "Mathew")
	contact("Cecilia Chapman,711-2880 Nulla ,St.Mankato Mississippi", 5567898800)
	fmt.Println("Added: ", add(45, 60))
	fmt.Println("Subtracted: ", sub(80, 60))

}
