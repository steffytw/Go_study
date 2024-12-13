package main

import (
	"fmt"
	"sync"
)

func RecieveChannel(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	x := <-ch
	fmt.Printf("x: %v\n", x)
}

func SendChannel(ch chan<- int, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- value
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	wg.Add(2)
	go SendChannel(ch, 100, &wg) // Sending data to the channel
	go RecieveChannel(ch, &wg)   // Receiving data from the channel
	wg.Wait()

}
