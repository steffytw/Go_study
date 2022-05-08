package main

import (
	"fmt"
)

func main() {

	// for loop is the only loop available in Go

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
		The Range Keyword:

			The range keyword is used to more easily iterate over an array, slice or map.

			It returns both the index and the value.

			The range keyword is used like this:

			Syntax:

				for index, value := array|slice|map {
				// code to be executed for each iteration
				}

	*/

	fruits := [3]string{"apple", "orange", "banana"}

	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
