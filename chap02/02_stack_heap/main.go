package main

import "fmt"

type user struct {
	name  string
	email string
}

func stayOnStack() user {
	u := user{
		name:  "Bill",
		email: "bill@email.com",
	}
	return u
}

func escapeToHeap() *user {
	u := &user{
		name:  "Bill",
		email: "bill@email.com",
	}
	return u
}

func main() {
	fmt.Println("stay on stack:", stayOnStack())
	fmt.Println("escape to heap:", escapeToHeap())
}
