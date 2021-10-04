package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct{}

func (Car) String() string {
	return "Car type"
}

type Cloud struct{}

func (Cloud) String() string {
	return "Cloud type"
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mvs := []fmt.Stringer{Car{}, Cloud{}}

	for i := 0; i < 10; i++ {
		rn := rand.Intn(2)
		if v, ok := mvs[rn].(Cloud); ok {
			fmt.Printf("Got lucky! -> %v\n", v)
		}
		fmt.Printf("Unlucky better luck next time..\n")
	}
}
