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
	var person1 Person

	person2 := Person{name: "Raveena", age: 35, job: "Actress", salary: 1000000}

	person1.name = "Ravi"
	person1.age = 23
	person1.job = "Teacher"
	person1.salary = 4000

	fmt.Println(person1.name) //Ravi

	// Print Person1 info by calling a function
	printPerson(person1)
	printPerson(person2)

}
