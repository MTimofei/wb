// Реализовать структуру-счетчик,
// которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count uint64
	mu    *sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		count: 0,
		mu:    &sync.Mutex{},
	}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

var c = NewCounter()
var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(*Counter) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				c.Inc()
			}
		}(c)
	}
	wg.Wait()
	fmt.Println(c.Value())
}
