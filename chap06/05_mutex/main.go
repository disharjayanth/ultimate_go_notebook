package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	const grs = 2
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(grs)

	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {
				mu.Lock()
				value := counter
				value++
				counter = value
				mu.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
