package main

import (
	"errors"
	"fmt"
)

func webCall() error {
	return errors.New("bad request")
}

func main() {
	if err := webCall(); err != nil {
		// Go's fmt package calls .Error method on err and formats to string
		fmt.Printf("%T\t%v\n", err, err)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Life is good!")
}
