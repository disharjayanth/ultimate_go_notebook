package main

import "fmt"

type customError struct{}

func (c *customError) Error() string {
	return "Find the bug."
}

func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error
	if _, err = fail(); err != nil {
		fmt.Println("why did this fail!")
		return
	}
	fmt.Println("No error!")
}
