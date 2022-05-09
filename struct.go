package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var p Person

	p.name = "Ravi"
	p.age = 23
	p.job = "Teacher"
	p.salary = 4000

	fmt.Println(p.name) //Ravi
}
