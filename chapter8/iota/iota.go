package main

import "fmt"

// iota is a function that simply increments by 1 each time, in the const block

const (
	first  = 1 << iota // 1 << 0 -> 1
	second             // 1 << 1 -> 2
	third              // 1 << 2 -> 4
	fourth             // 1 << 3 -> 8
	fifth              // 1 << 4 -> 16
)

func main() {
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(fifth)
}