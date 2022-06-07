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
	fmt.Printf("Sending email to %s:%s\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	joe := admin{
		user:  user{name: "Joe", email: "joe@admins.org"},
		level: "manager",
	}

	// this works since the inner type's methods are promoted to the outer type
	sendNotification(&joe)
}