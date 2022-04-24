package main

import "fmt"

func main() {
	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else { // brackets in the else statement should be like } else {: otherwise it raise errors
		fmt.Println("It is cold out there")
	}
}
