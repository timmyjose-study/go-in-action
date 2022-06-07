package main

import (
	"entities1"
	"fmt"
)

func main() {
	joe := entities1.Admin{Level: "supervisor"}

	// this still works since the package-private inner type's
	// fields are promoted to the outer type, though the inner
	// type cannot be used directly
	joe.Name = "Joe"
	joe.Email = "joe@samizdat.org"

	fmt.Printf("%v\n", joe)
}
