package main

/*

Unidirectional Channels

Unidirectional channels restrict the direction of data flow, allowing only sending or
receiving operations. They are declared using the arrow (<-) notation to indicate the
data flow direction.

*/

func SendChannel(ch chan<- int) {
	ch <- 100
}

func main() {
	//send only channel
	ch := make(chan<- int) // Declaring a send-only channel
	go SendChannel(ch)
}
