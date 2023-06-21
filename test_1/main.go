package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go incrementCounter(&wg)
	go incrementCounter(&wg)

	wg.Wait()

	fmt.Println("Final counter value:", counter)
}

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		// Read the current value of the counter
		current := counter

		// Increment the counter
		current++

		// Store the updated value back to the counter
		counter = current
	}
}
