package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func displayUser(u *user) {
	fmt.Printf("Name: %s, email: %s, extension: %d, privileged? %v\n", u.name, u.email, u.ext, u.privileged)
}

type admin struct {
	person user
	level  string
}

func main() {
	var bill user
	displayUser(&bill)

	lisa := user{name: "Lisa", email: "lisa@lisa.com", ext: 199, privileged: true}
	displayUser(&lisa)

	fred := admin{
		person: user{
			name:       "Fred",
			email:      "fred@bob.com",
			ext:        1,
			privileged: true,
		},
		level: "admin",
	}

	displayAdmin(&fred)

}

func displayAdmin(a *admin) {
	fmt.Printf("user details: %v, level: %s\n", a.person, a.level)
}