// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
// с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var list = []int{2, 4, 6, 8, 10}

	first(list)
}

func first(list []int) {
	var ch = make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, n := range list {
			ch <- n * n
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var sum int
		for n := range ch {
			sum += n
		}
		fmt.Println(sum)
	}()

	wg.Wait()
}
