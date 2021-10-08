package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// sets max no. of cpus that can be used for executing simultaneously
// func init() {
// 	runtime.GOMAXPROCS(1)
// }

func upperCase() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
}

func lowerCase() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
}

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	fmt.Println("Number of CPUs executing simultaneously", runtime.NumCPU())

	wg.Add(2)

	go func() {
		upperCase()
		wg.Done()
	}()

	go func() {
		lowerCase()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Ended:", time.Since(now))
}
