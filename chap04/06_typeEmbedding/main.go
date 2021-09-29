package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	*user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	a := admin{
		user: &user{
			name:  "Bill",
			email: "bill@email.com",
		},
		level: "super",
	}

	a.user.notify()
	a.notify()

	sendNotification(&a)
}
