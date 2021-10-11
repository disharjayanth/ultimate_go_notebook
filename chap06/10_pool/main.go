package main

import (
	"fmt"
	"runtime"
	"time"
)

func pooling() {
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	fmt.Println("Max number of O.S threads:", g)

	for i := 0; i < g; i++ {
		go func(child int) {
			for str := range ch {
				fmt.Printf("child %d: received signal: %s\n", child, str)
			}
			fmt.Printf("child %d: receive signal shutdown\n", child)
		}(i)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data" + fmt.Sprintf(" %d", w)
		fmt.Printf("parent %d: sent signal\n", w)
	}

	close(ch)
	fmt.Println("parent:sent shutdown signal")

	time.Sleep(1 * time.Second)
	fmt.Println("---------------------------------")
}

func main() {
	pooling()
}
