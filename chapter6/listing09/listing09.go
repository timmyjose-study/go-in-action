// race condition example
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Printf("Final counter value = %d\n", counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		// yield the thread
		runtime.Gosched()
		value++
		counter = value
	}
}