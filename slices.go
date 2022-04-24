package main

import "fmt"

func main() {
	/*
		Slices:

			They are similar to arrays, but are more powerful and flexible.

			Slices are also used to store multiple values
			of the same type in a single variable.

			However, unlike arrays, the length of a slice can grow and shrink as you see fit.
	*/
	// var slice_1 = []int{1, 2, 3, 4, 5}
	slice_1 := []int{1, 2, 3, 4, 5}

	fmt.Println(slice_1) // [1 2 3 4 5]

	// len() function - returns the number of elements in the slice .
	fmt.Println(len(slice_1)) // 5

	//cap() function - returns the number of elements in which the slice can grow or shrink to of.
	fmt.Println(cap(slice_1)) // 5

	slice_2 := []string{"Slices", "are", "powerful", "and", "flexible"}
	fmt.Println(slice_2)
	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))

	//Create a Slice From an Array:

	myarray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // An array
	myslice := myarray[1:7]                          // A slice made from the array

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// make() function can also be used to create a slice
}
