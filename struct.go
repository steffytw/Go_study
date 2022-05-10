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

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}

func main() {

	/*
				Go Structures:
					A struct (short for structure) is used to create a
		 			collection of members of different data types, into a single variable.
	*/
	var p Person

	p.name = "Ravi"
	p.age = 23
	p.job = "Teacher"
	p.salary = 4000

	fmt.Println(p.name) //Ravi

	// Print Pers1 info by calling a function
	printPerson(p)
}
