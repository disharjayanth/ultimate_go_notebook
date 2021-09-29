package main

import "fmt"

type printer interface {
	print()
}

type cannon struct {
	name string
}

func (c cannon) print() {
	fmt.Printf("Printer Name: %s\n", c.name)
}

type espon struct {
	name string
}

func (e *espon) print() {
	fmt.Printf("Printer Name: %s\n", e.name)
}

func main() {
	c := cannon{
		name: "PIXMA TR4520",
	}

	e := &espon{
		name: "Workforce Pro WF-3720",
	}

	printers := []printer{c, e}

	c.name = "PRO-1000"
	e.name = "XP-4100"

	for _, p := range printers {
		p.print()
	}

	fmt.Printf("%v\n%v\n", c, *e)
}
