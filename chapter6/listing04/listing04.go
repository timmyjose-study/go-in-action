package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Creating goroutines")
	go printPrimes("A")
	go printPrimes("B")

	wg.Wait()
	fmt.Println("All done!")
}

func printPrimes(prefix string) {
	defer wg.Done()

next:
	for n := 2; n <= 5000; n++ {
		for d := 2; d < n; d++ {
			if n%d == 0 {
				continue next
			}
		}

		fmt.Printf("[%s] %d\n", prefix, n)
	}

	fmt.Println("Completed")
}