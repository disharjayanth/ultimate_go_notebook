package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	users := make(map[string]user)

	users["John"] = user{
		name: "John",
		age:  22,
	}

	users["Joee"] = user{
		name: "Joe",
		age:  20,
	}

	for key, value := range users {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}

	value, exists := users["John"]
	fmt.Println("Value: ", value, "Exists:", exists)

	value, exists = users["Rogan"]
	fmt.Println("Value: ", value, "Exists:", exists)

	delete(users, "Joee")
	fmt.Println("users map after deleting Joee:", users)
}
