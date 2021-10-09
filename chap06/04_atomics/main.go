package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	const grs = 2
	var counter int32
	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for j := 0; j < 2; j++ {
				atomic.AddInt32(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
