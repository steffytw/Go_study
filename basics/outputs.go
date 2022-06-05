package main

import "fmt"

func main() {
	a, b := "Hello", "World"
	var i, j = 10, 20.00

	/*
		Print() Function:

			The Print() function prints its arguments with their default format.
	*/

	fmt.Print(a)
	fmt.Print(b)

	fmt.Print(a, "\n") //print the arguments in new lines, we need to use \n.
	fmt.Print(b, "\n")

	fmt.Print(a, "\n", b, "\n")

	fmt.Print(a, " ", b, "\n")

	fmt.Print(i, j) // inserts a space between the arguments if neither are strings
	fmt.Print("\n")

	/*
		Println() Function :

			The Println() function is similar to Print() with the difference
			that a whitespace is added between the arguments, and a newline is added at the end.
	*/

	fmt.Println("Println : ", a, b)

	/*
		Printf() Function:

			The Printf() function first formats its argument based on the given formatting
			verb and then prints them. Here we will use two formatting verbs:

				%v is used to print the value of the arguments
				%T is used to print the type of the arguments
	*/

	fmt.Printf("Value of i is %v and Type of i is %T \n", i, i)
	fmt.Printf("Value of j is %v and Type of j is %T \n", j, j)

}
