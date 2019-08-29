package main

import (
	"fmt"
	"sync"
)

// Global variable
var counter int

func main() {
	const grs = 2

	// WaitGroup of 2 go routines, after all go routines call Done() then Wait() release
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create 2 go routines
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				// get from memory
				value := counter

				// increase
				value++

				fmt.Println(value)

				// set to memory
				counter = value
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
