package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}

func start() {
	go start2()
	fmt.Println("In Goroutine")
}
func start2() {
	fmt.Println("In Goroutine2")
}
