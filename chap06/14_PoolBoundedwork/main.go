package main

import (
	"fmt"
	"runtime"
	"sync"
)

func boundedWorkPool() {
	work := []string{"paper", "paper", "paper", 2000: "paper"}

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for i := 0; i < g; i++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d: recevied signal: %s\n", child, wrk)
			}
			fmt.Printf("child %d: signal shutdown signal\n", child)
		}(i)
	}

	for _, wrk := range work {
		ch <- wrk
	}

	close(ch)
	wg.Wait()

	fmt.Println("---------------------------------------------------------------------------")
}

func main() {
	boundedWorkPool()
}
