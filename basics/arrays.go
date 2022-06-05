package main

import "fmt"

func main() {
	/*
		Arrays:
			They are used to store multiple values of the same type
			in a single variable, instead of declaring separate variables for each value.
	*/

	// var array_1 = [5]int{1, 2, 3, 4, 5}
	array_1 := [...]int{1, 2, 3, 4, 5}     // full initialized
	array_2 := [5]int{11}                  // partially initialized
	array_3 := [5]int{}                    // not initialized
	array_4 := [5]int{1: 10, 2: 40, 3: 50} // initialize with index:value

	fmt.Println(array_1) // [1 2 3 4 5]

	array_1[0] = 10
	fmt.Println(array_1) // [10 2 3 4 5]

	fmt.Println(array_1[0])   // 10
	fmt.Println(array_2)      // [11 0 0 0 0]
	fmt.Println(array_3)      // [0 0 0 0 0]
	fmt.Println(array_4)      // [0 10 40 50 0]
	fmt.Println(len(array_4)) // 5
	fmt.Println(cap(array_4)) // 5
}
