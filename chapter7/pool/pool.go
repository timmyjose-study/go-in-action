package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	mu        sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("pool is closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size is too small to create a pool")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil

}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "Shared resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire: ", "New resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release: ", "In queue")
	default:
		log.Println("Release: ", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}