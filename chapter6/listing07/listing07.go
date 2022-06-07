package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for c := 'a'; c <= 'a'+25; c++ {
				fmt.Printf("%c ", c)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for c := 'A'; c <= 'A'+25; c++ {
				fmt.Printf("%c ", c)
			}
		}
	}()

	wg.Wait()

	fmt.Println("All done!")
}