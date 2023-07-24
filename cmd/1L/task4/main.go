// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout
// Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
)

type Data struct{}

// при старет указываем флаг WORKER с нужным количеством
var w = flag.Int("WORKER", runtime.NumGoroutine(), "number of workers")

var stream []Data
var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	flag.Parse()
	fmt.Println(*w)
	var mainCh = make(chan Data, *w)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, volume := range stream {
			mainCh <- volume
		}
	}()

	for i := 0; i < *w; i++ {
		go worker(mainCh, wg)
	}

	wg.Add(1)
	// при нажаии Ctrl+C будит закрываться main chan
	go func() {
		defer wg.Done()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		close(mainCh)
	}()

	// функция main не завершится пока не отработают все worker
	wg.Wait()
}

// покак канал открыт, worker будеьт работать
func worker(mainCh chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range mainCh {
		fmt.Println(v)
	}
}
