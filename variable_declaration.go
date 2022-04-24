package main

import (
	"fmt"
)

func main() {
	var name string = "John" //type is string
	var car = "BMW"          //type is inferred
	x := 2                   //type is inferred

	fmt.Println(name)
	fmt.Println(car)
	fmt.Println(x)
}
