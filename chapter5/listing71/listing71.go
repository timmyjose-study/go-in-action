package main

import "fmt"
import "entities"

func main() {
	bob := entities.User{Name: "Bob", Email: "bob@users.org"}
	fmt.Printf("%v\n", bob)
}
