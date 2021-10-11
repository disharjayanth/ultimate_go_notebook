package main

import (
	"fmt"
	"sync"
	"time"
)

var data []string
var rwMutex sync.RWMutex

func writer(i int) {
	rwMutex.Lock()
	{
		time.Sleep(1 * time.Second)
		fmt.Println("********> Performing write")
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
}

func reader(id int) {
	rwMutex.RLock()
	{
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d: Performing Read: Length[%d]\n", id, len(data))
	}
	rwMutex.RUnlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			writer(i)
		}
		wg.Done()
	}()

	for i := 0; i < 8; i++ {
		go func(id int) {
			for {
				reader(id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Complete!")
}
