package main

import "fmt"

type user struct {
	name  string
	email string
}

// Prints user name with respective user email.
func (u user) notify() {
	fmt.Printf("Sending User email to %s<%s>\n", u.name, u.email)
}

// Accepts email which is to be changed.
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
