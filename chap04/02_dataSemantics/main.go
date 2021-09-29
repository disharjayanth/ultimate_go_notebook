package main

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Printf("My name is: %s age is: %d \n", d.name, d.age)
}

func (d *data) setAge(age int) {
	d.age = age
}

func main() {
	d := data{
		name: "John",
		age:  28,
	}

	f1 := d.displayName
	f1()
	d.name = "Bill"
	f1()

	f2 := d.setAge
	f2(30)
	d.name = "Sammy"
	f2(42)
	d.displayName()
}
