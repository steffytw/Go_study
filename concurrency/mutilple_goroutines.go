package main

import (
	"fmt"
	"time"
)

func execute(id int) {
	fmt.Printf("id: %d\n", id)
}

func main() {
	fmt.Println("Started")
	for i := 0; i < 10; i++ {
		go execute(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Finished")
}
