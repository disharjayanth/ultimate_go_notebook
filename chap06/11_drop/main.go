package main

import (
	"fmt"
	"time"
)

func drop() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for str := range ch {
			fmt.Printf("child: received signal: %s\n", str)
		}
	}()

	for i := 0; i < 2000; i++ {
		select {
		case ch <- "data":
			fmt.Printf("parent: sent signal: %d\n", i)
		default:
			fmt.Printf("parent: dropped data: %d\n", i)
		}
	}

	close(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("parent: shutdown signal.....")
}

func main() {
	drop()
}
