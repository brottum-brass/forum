package middleware

import (
	"net/http"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		count: 0,
	}
}

func (c *Counter) Next(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.increment()

		next.ServeHTTP(w, r)
	})
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func (c *Counter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count = 0
}
