package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s:%s\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s:%s\n", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	bob := user{name: "Bob", email: "bob@bobby.org"}
	sendNotification(&bob)

	joe := admin{name: "Joe", email: "jo@admins.org"}
	sendNotification(&joe)
}