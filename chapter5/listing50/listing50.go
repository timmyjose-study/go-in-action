package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending email to %s:%s\n", u.name, u.email)
}

type admin struct {
	user  // embedded or inner type
	level string
}

func main() {
	joe := admin{
		user: user{
			name:  "Joe",
			email: "joe@admins.org",
		},
		level: "manager",
	}

	// methods on the inner type are available
	joe.user.notify()

	// the inner type's methods are also available directly
	// on the outer type
	joe.notify()
}