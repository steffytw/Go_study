package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2) // Creating a buffered channel with a capacity of 2

	ch <- 1 // Sending the value 1 to the channel
	ch <- 2 // Sending the value 2 to the channel

	x := <-ch      // Receiving the value from the channel
	fmt.Println(x) // Output: 1

	y := <-ch      // Receiving the value from the channel
	fmt.Println(y) // Output: 2
}
