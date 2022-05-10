package main

import (
	"fmt"
)

func main() {

	// 	Maps are used to store data values in key:value pairs.

	// A map is an unordered and changeable collection that does not allow duplicates.

	// The default value of a map is nil.

	// Maps hold references to an underlying hash table.

	// Go has multiple ways for creating maps.

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	c := make((map[int]string))
	c[1] = "apple"
	c[2] = "orange"
	c[3] = "kiwi"

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	fmt.Printf("b\t%v\n", c)
}
