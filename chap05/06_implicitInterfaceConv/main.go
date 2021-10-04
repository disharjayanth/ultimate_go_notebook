package main

import "fmt"

type bike struct{}

func (bike) Move() {
	fmt.Println("Move the bike")
}

func (bike) Lock() {
	fmt.Println("Lock the bike")
}

func (b bike) Unlock() {
	fmt.Println("Unlock the bike")
}

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

var ml MoveLocker
var m Mover

func main() {
	ml = bike{}
	m = ml
	fmt.Printf("%v\t%T\n", m, m)

	// // ml cannot be equal to m since ml of type MoveLocker, m is of type Mover
	// // when m = ml, complier knows ml(MoveLocker) implements Mover interface
	// // but ml = m, gives error because compiler knows m only implements Mover not MoveLocker
	// ml = m
}
