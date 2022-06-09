package main

import (
	"gia/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (conn *dbConnection) Close() error {
	log.Println("Close: connection", conn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New connection", id)
	return &dbConnection{id}, nil

}

func performQuery(query int, pool *pool.Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	pool, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for q := 0; q < maxGoroutines; q++ {
		go func(q int) {
			performQuery(q, pool)
			wg.Done()
		}(q)
	}

	wg.Wait()

	log.Println("Shutdown program")
	pool.Close()
}