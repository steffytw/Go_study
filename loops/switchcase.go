package main

import (
	"fmt"
)

func main() {
	day := 7

	switch day {
	case 1: //All the case values should have the same type as the switch expression
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default: //run if there is no case match
		fmt.Println("Not a weekday")
	}
}
