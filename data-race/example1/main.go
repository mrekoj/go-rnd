package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				value := counter
				value++
				counter = value
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
