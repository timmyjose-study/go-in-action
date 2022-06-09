package work

import "sync"

type Worker interface {
	Work()
}

type Pool struct {
	tasks chan Worker
	wg    sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		tasks: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.tasks {
				w.Work()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker) {
	p.tasks <- w
}

func (p *Pool) Shutdown() {
	close(p.tasks)
	p.wg.Wait()
}