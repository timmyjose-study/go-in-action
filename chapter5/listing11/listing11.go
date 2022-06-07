package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending email to %s:%s\n", u.name, u.email)
}

func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}

func main() {
	bob := user{name: "Bob", email: "bobber@bobbers.com"}
	bob.notify()

	bob.changeEmail("bob@bobbers.com")
	bob.notify()

	lisa := &user{name: "Lisa", email: "lisa@example.com"}
	lisa.notify()

	lisa.changeEmail("lisa@lisa.com")
	lisa.notify()
}