package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func fanOutInSem() {
	children := 2000
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- "data" + fmt.Sprintf(" %d", child)
				fmt.Println("child: sent signal:", child)
			}
			<-sem
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println("parent: received signal:", d)
	}

	fmt.Println("****************************")
}

func main() {
	fanOutInSem()
}
