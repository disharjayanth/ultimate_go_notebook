package main

import (
	"fmt"
	"time"
)

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "some work"
		fmt.Println("child: sent signal")
	}()

	d, ok := <-ch
	fmt.Println("Parent received signal:", d, ok)

	time.Sleep(1 * time.Second)
	fmt.Println("-------------------------------")
}

func main() {
	waitForResult()
}
