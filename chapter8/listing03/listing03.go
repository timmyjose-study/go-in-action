package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// default is to Stdout
	log.Println("message")

	// Println followed by os.Exit(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		log.Fatalln("fatal message")
	}()

	// Println followed by panic
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		log.Panicln("panic message")
	}()

	wg.Wait()
}