// fixing the race condition using a mutex
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mu      sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Printf("Final value of counter: %d\n", counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mu.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mu.Unlock()
	}
}