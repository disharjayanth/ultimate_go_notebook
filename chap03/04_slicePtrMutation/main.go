package main

import (
	"encoding/binary"
	"fmt"
)

type user struct {
	likes int
}

func main() {
	users := make([]user, 1)
	ptrUser0 := &users[0]
	ptrUser0.likes++

	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	users = append(users, user{})
	ptrUser0.likes++

	fmt.Println("*************")
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	fmt.Println("ptrUser0:", ptrUser0)

	fmt.Println("Linear Traversal")

	x := []byte{0x0A, 0x15, 0x0e, 0x28, 0x05, 0x96, 0x0b, 0xd0, 0x0}
	a := x[0]
	b := binary.LittleEndian.Uint16(x[1:3])
	c := binary.LittleEndian.Uint16(x[3:5])
	d := binary.LittleEndian.Uint16(x[5:9])

	fmt.Println(a, b, c, d)
}
