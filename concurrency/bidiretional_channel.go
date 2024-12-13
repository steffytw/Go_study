package main

import (
	"fmt"
	"time"
	// concurrency "/usr/local/go/src/github.com/steffytw/Go_study/concurrency/test.go"
)

func RecieveChannel(ch <-chan int) {
	x := <-ch
	fmt.Printf("x: %v\n", x)
}

func SendChannel(ch chan<- int) {
	ch <- 100
}

func main() {

	ch := make(chan int, 1)
	go SendChannel(ch) // Sending data to the channel
	// go concurrency.RecieveChannel(ch) // Receiving data from the channel
	go RecieveChannel(ch)
	time.Sleep(time.Second) // Wait for goroutines to complete
}
