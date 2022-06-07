package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	baton := make(chan int)

	wg.Add(1)
	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runenr %d running with the baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d finished, Race over\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d exchanges baton with runner %d\n", runner, newRunner)
	baton <- newRunner
}