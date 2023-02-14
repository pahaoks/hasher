package cacher

import (
	"context"
	"sync"
	"time"
)

// Incapsulation of cacher logic for future reuse
type MemoryCacher[R any] struct {
	lock sync.Mutex

	interval time.Duration
	provide  func() *R
	result   *R
}

// Instantiate memory cacher
func NewMemoryCacher[R any](
	interval time.Duration,
	provide func() *R,
) *MemoryCacher[R] {
	return &MemoryCacher[R]{
		lock:     sync.Mutex{},
		interval: interval,
		provide:  provide,
		result:   nil,
	}
}

// Return last cached result
func (c *MemoryCacher[R]) GetResult() *R {
	c.lock.Lock()

	if c.result == nil {
		c.result = c.provide()

		go c.Start(context.Background())
	}

	result := c.result

	c.lock.Unlock()

	return result
}

// Start updating cache according to intervals
func (c *MemoryCacher[R]) Start(ctx context.Context) {
	for {
		time.Sleep(c.interval)

		result := c.provide()

		if result != nil {
			c.lock.Lock()
			c.result = result
			c.lock.Unlock()
		}
	}
}
