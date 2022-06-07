package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	numTasks   = 10
	numWorkers = 4
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	tasks := make(chan string, numTasks)

	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(tasks, i)
	}

	for i := 0; i < numTasks; i++ {
		tasks <- fmt.Sprintf("Task %d", i)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks <-chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d shutting down...\n", worker)
			return
		}

		fmt.Printf("Worker %d starting task %s\n", worker, task)

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		fmt.Printf("Worker %d finished task %s\n", worker, task)
	}
}