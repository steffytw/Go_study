package concurrency

import "fmt"

func RecieveChannel(ch <-chan int) {
	x := <-ch
	fmt.Printf("x: %v\n", x)
}

func SendChannel(ch chan<- int) {
	ch <- 100
}
