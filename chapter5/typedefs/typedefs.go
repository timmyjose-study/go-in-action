package main

import "fmt"

type Duration int64

func displayDuration(d Duration) {
	fmt.Printf("%d\n", d)
}

func main() {
	d := Duration(1000)
	displayDuration(d)
}