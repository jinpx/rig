package u_cycle

import (
	"sync"
	"sync/atomic"
)

// TCycle ..
type TCycle struct {
	mu      *sync.Mutex
	wg      *sync.WaitGroup
	done    chan struct{}
	quit    chan error
	closing uint32
	waiting uint32
	// works []func() error
}

// NewCycle new a cycle life
func NewCycle() *TCycle {
	return &TCycle{
		mu:      &sync.Mutex{},
		wg:      &sync.WaitGroup{},
		done:    make(chan struct{}),
		quit:    make(chan error),
		closing: 0,
		waiting: 0,
	}
}

// Run a new goroutine
func (c *TCycle) Run(fn func() error) {
	c.mu.Lock()
	//todo add check options panic before waiting
	defer c.mu.Unlock()
	c.wg.Add(1)
	go func(c *TCycle) {
		defer c.wg.Done()
		if err := fn(); err != nil {
			c.quit <- err
		}
	}(c)
}

// Done block and return a chan error
func (c *TCycle) Done() <-chan struct{} {
	if atomic.CompareAndSwapUint32(&c.waiting, 0, 1) {
		go func(c *TCycle) {
			c.mu.Lock()
			defer c.mu.Unlock()
			c.wg.Wait()
			close(c.done)
		}(c)
	}
	return c.done
}

// DoneAndClose ..
func (c *TCycle) DoneAndClose() {
	<-c.Done()
	c.Close()
}

// Close ..
func (c *TCycle) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if atomic.CompareAndSwapUint32(&c.closing, 0, 1) {
		close(c.quit)
	}
}

// Wait blocked for a life cycle
func (c *TCycle) Wait() <-chan error {
	return c.quit
}
