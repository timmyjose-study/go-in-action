package main

import (
	"counters"
	"fmt"
)

func main() {
	counter := counters.New(10)
	fmt.Println(counter)
}
