package main

import (
	"fmt"

	"github.com/disharjayanth/ultimate_go_notebook/chap04/07_export/counters"
	"github.com/disharjayanth/ultimate_go_notebook/chap04/07_export/users"
)

func main() {
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)

	m := users.Manager{
		Title: "Go Programmer",
	}

	fmt.Println("Manger:", m)
}
