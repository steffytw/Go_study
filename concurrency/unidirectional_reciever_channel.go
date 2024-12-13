package main

import "fmt"

/*

Unidirectional Channels

Unidirectional channels restrict the direction of data flow, allowing only sending or
receiving operations. They are declared using the arrow (<-) notation to indicate the
data flow direction.

*/

func RecieveChannel(ch <-chan int) {
	x := <-ch
	fmt.Printf("x: %v\n", x)
}

func main() {

	ch := make(<-chan int) // Declaring a receive-only channel
	go RecieveChannel(ch)
}
