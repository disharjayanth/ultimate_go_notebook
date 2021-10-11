package main

import (
	"fmt"
	"math/rand"
	"time"
)

func waitForTask() {
	ch := make(chan string)

	go func() {
		d := <-ch
		fmt.Println("child: received signal:", d)
	}()

	fmt.Println("some work....")
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "data"
	fmt.Println("parent: sent signal")

	time.Sleep(1 * time.Second)
}

func main() {
	waitForTask()
}
