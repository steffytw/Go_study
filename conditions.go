package main

import "fmt"

func main() {
	/*
		Go Conditions:

		A condition can be either true or false.

		Go supports the usual comparison operators from mathematics:

			- Less than <
			- Less than or equal <=
			- Greater than >
			- Greater than or equal >=
			- Equal to ==
			- Not equal to !=

		Additionally, Go supports the usual logical operators:

			- Logical AND &&
			- Logical OR ||
			- Logical NOT !

	*/

	// if statement

	a := "hello"

	if a == "hello" {
		fmt.Printf("%v!\n", a)
	}

	// if else statement
	temperature := 14

	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else { // brackets in the else statement should be like } else {: otherwise it raise errors
		fmt.Println("It is cold out there")
	}

	//else if statement

	val1 := 100
	val2 := 78
	if val1 < val2 {
		fmt.Printf("%v is less than %v.\n", val1, val2)
	} else if val1 > val2 {
		fmt.Printf("%v is greater than %v.\n", val1, val2)
	} else {
		fmt.Println("Both are equal.")
	}

	//nested if

	num := 20
	if num >= 10 {
		fmt.Printf("%v is more than 10.\n", num)
		if num > 15 {
			fmt.Printf("%v is also more than 15.\n", num)
		}
	} else {
		fmt.Printf("%v is less than 10.\n", num)
	}

	// Logical OR ||

	protocol_used := "TCP"
	if protocol_used == "TCP" || protocol_used == "UDP" {
		fmt.Printf("Protocol used is %v\n", protocol_used)
	}

}
