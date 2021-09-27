package main

import "fmt"

func main() {
	strings := [5]string{"Apple", "Banana", "Pumpkin", "Pinneapple", "Orange"}

	for i, v := range strings {
		fmt.Println("Value semantic:", i, v)
	}

	for i := range strings {
		fmt.Println("Pointer semantic:", strings[i])
	}

	type person struct {
		name string
		age  int
	}

	p := person{
		name: "John",
		age:  22,
	}

	interfaceType := []interface{}{"string", true, 2, p}

	for _, value := range interfaceType {
		fmt.Printf("%v\t%T\n", value, value)
	}
}
