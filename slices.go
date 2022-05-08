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
	fmt.Println(slice_2)      //[Slices are powerful and flexible]
	fmt.Println(len(slice_2)) //5
	fmt.Println(cap(slice_2)) //5

	//Create a Slice From an Array:

	myarray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // An array
	myslice := myarray[1:7]                          // A slice made from the array

	fmt.Printf("myslice = %v\n", myslice)       //myslice = [1 2 3 4 5 6]
	fmt.Printf("length = %d\n", len(myslice))   // length = 6
	fmt.Printf("capacity = %d\n", cap(myslice)) //capacity = 9

	/*
		make() function can also be used to create a slice

		slice_name := make([]type, length, capacity)
	*/

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	myslice1[1] = 12
	fmt.Printf("Modified myslice1 = %d\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

}
