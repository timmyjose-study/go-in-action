package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // only one logical processor

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for c := 'a'; c <= 'a'+25; c++ {
			fmt.Printf("%c ", c)
		}
		fmt.Println()
	}()

	go func() {
		defer wg.Done()

		for c := 'A'; c <= 'A'+25; c++ {
			fmt.Printf("%c ", c)
		}
		fmt.Println()
	}()

	wg.Wait()
	fmt.Println("All done!")
}