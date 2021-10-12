package main

import (
	"context"
	"fmt"
	"time"
)

func cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		// Simulatation of some work
		time.Sleep(time.Duration(152) * time.Millisecond)
		ch <- "data"
		close(ch)
	}()

	select {
	case data := <-ch:
		fmt.Println("Received data from ch channel:", data)
	case <-ctx.Done():
		fmt.Println("Timeout!, Required time to complete surpassed!")
	}
}

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		cancellation()
	}
}
