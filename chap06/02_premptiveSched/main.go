package main

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"sync"
)

func printHashes(prefix string) {
	for i := 1; i <= 50000; i++ {
		num := strconv.Itoa(i)
		sum := sha1.Sum([]byte(num))
		fmt.Printf("%s: %5d: %x\n", prefix, i, sum)
	}
	fmt.Println("Completed", prefix)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		printHashes("A")
		wg.Done()
	}()

	go func() {
		printHashes("B")
		wg.Done()
	}()

	fmt.Println("Waiting to finish.")
	wg.Wait()

	fmt.Println("Terminating program!")
}
