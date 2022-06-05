package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) details() {
	fmt.Printf("Name: %s\n", e.name)
	fmt.Printf("Age: %d\n", e.age)

}

func (e *employee) setName(newName string) {
	e.name = newName
	return
}
func (e employee) setAge(newAge int) {
	e.age = newAge
	return
}

func (e employee) getSalary() int {
	return e.salary
}

func main() {
	emp := employee{name: "Sam", age: 31, salary: 200000}
	emp.details()

	emp.setName("Johns")
	fmt.Printf("Updated Name: %s\n", emp.name)

	(&emp).setAge(40)
	fmt.Printf("Updated Age: %s\n", emp.age)

	emp.details()

	fmt.Printf("Salary %d\n", emp.getSalary())
}
