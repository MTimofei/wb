// Реализовать конкурентную запись данных в map.

package main

import (
	"sync"
)

var m = make(map[int]int)
var ms = sync.Map{}
var mu = &sync.Mutex{}
var n int

func main() {

	// конкурентная запись в мап из пакета sync
	go func() { ms.Store(n, n) }()

	// конкурентная запись в стандартную мап с помощью мьютекса
	go func(int) {
		mu.Lock()
		defer mu.Unlock()
		m[n] = n + 1
	}(n)

}
