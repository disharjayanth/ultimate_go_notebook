package main

import "fmt"

// Grouping different concrete type with common behavior using interface Speaker.
type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Println("Woof! My name is ", d.Name, "it is", d.IsMammal, "i am a mammal with pack factor of", d.PackFactor)
}

type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Println("Meow! My name is ", c.Name, "it is", c.IsMammal, "i am a mammal with pack factor of", c.ClimbFactor)
}

func main() {
	speakers := []Speaker{
		&Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},
		&Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	for _, speaker := range speakers {
		speaker.Speak()
	}
}
