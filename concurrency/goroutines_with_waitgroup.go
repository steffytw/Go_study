package main

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am first runner")
}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am second runner")
}

func execute(wg *sync.WaitGroup) {

	go runner1(wg)
	go runner2(wg)

}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	execute(&wg)
	wg.Wait()
	// time.Sleep(time.Second)
	fmt.Println("All runners finished")
}
