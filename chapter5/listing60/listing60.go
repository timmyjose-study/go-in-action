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
	user
	level string
}

func (ad *admin) notify() {
	fmt.Printf("Sending admin email to %s:%s\n", ad.name, ad.email)
}

func sendNoiification(n notifier) {
	n.notify()
}

func main() {
	bob := user{name: "Bob", email: "bob@users.com"}
	sendNoiification(&bob)

	joe := admin{user: user{name: "Joe", email: "joe@admins.com"}, level: "manager"}
	sendNoiification(&joe)

	// this works as expected
	joe.user.notify()

	// this does not work on the innner type since `admin` has its own implementation of `notify`
	joe.notify()
}