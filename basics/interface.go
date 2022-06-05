package main

import (
	"fmt"
)

type animal interface {
	breathe()
	walk()
}
type dog struct {
	age int32
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walks")
}
func main() {
	var a animal
	a = dog{age: 10}
	a.breathe() //Dog breathes"
	a.walk()    //Dog walks"
}
