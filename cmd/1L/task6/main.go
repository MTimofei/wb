// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"sync"
	"time"
)

func _() {
	// Выполнение кода горутины

	// Завершение выполнения горутины
	return
}

func _(ctx context.Context) {
	// Выполнение кода горутины

	select {
	case <-ctx.Done():
		// Горутина остановлена по причине отмены
		return
	case <-time.After(5 * time.Second):
		// Горутина остановлена по причине таймаута
		return
	}
}

func _(stop chan bool) {
	// Выполнение кода горутины

	select {
	case <-stop:
		// Горутина остановлена
		return
	}
}

func myGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()

	// Выполнение кода горутины
}

func main() {
	var wg sync.WaitGroup

	// Запустите горутину
	wg.Add(1)
	go myGoroutine(&wg)

	// Ожидание завершения горутины
	wg.Wait()
}
