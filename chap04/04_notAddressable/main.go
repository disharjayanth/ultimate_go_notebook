package main

import "fmt"

//constants are not addressable!
type duration int

func (d *duration) notify() {
	fmt.Printf("Sending notification in %d\n", *d)
}

func main() {
	d := duration(42)
	d.notify()
}
