package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	u := user{
		name:  "John",
		email: "john@email.com",
	}

	u.notify()

	u.changeEmail("john@xyz.com")

	fmt.Println("After email change")
	u.notify()
}
