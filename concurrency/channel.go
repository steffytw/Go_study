package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int)
	// ch := make(chan int)
	fmt.Println("Created channel : ", ch)

	fmt.Println("Sending the data to the channel")
	go send_channel(ch)

	fmt.Println("Receiving the data from the channel")
	go receive_channel(ch)

	time.Sleep(time.Second * 1)
}

func send_channel(ch chan int) {
	ch <- 1
}
func receive_channel(ch chan int) {
	val := <-ch
	fmt.Printf("Value Received=%d in receive function\n", val)
}
