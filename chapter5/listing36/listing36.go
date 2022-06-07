package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// Method sets:
// A receiver of type (t T) can receive either *T or T
// A receiver of type (t *T) can only receive *T
func (u *user) notify() {
	fmt.Printf("Sending email to %s:%s\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{name: "Bob", email: "bob@users.com"}
	u.notify()           // works
	sendNotification(&u) // also works
	// this doesn't work
	//sendNotification(u)
}