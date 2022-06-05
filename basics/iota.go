package main

import "fmt"

// const (
// 	a = iota
// 	b
// 	c
// )

// const ( // similar to the above constant
// 	a = iota
// 	b = iota
// 	c = iota
// )

const ( // siota expressions – iota allows expressions which can be used to set any value for the constant
	a = iota // o/p: 0 6 12
	b = iota + 5
	c = iota * 6
)

// const ( // iota increment can be skipped using a blank identifier
// 	a = iota // o/p: 0 2 3
//	_
// 	b
// 	c
// )

type Size uint8

const (
	small Size = iota
	medium
	large
	extraLarge
)

func main() {
	/*
		IOTA is

			- A counter which starts with zero
			- Increases by 1 after each line
			- Is only used with constant
	*/
	fmt.Println(a, b, c)

	//Enum in go

	fmt.Println(small)      //0
	fmt.Println(medium)     // 1
	fmt.Println(large)      //2
	fmt.Println(extraLarge) //3

	var m Size = 2
	m.toString() // o/p: Large

}

func (s Size) toString() {
	switch s {
	case small:
		fmt.Println("Small")
	case medium:
		fmt.Println("Medium")
	case large:
		fmt.Println("Large")
	case extraLarge:
		fmt.Println("Extra Large")
	default:
		fmt.Println("Invalid Size entry")
	}
}
