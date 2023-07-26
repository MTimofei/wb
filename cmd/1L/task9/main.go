// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
//  после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

var ch1 = make(chan int)
var ch2 = make(chan int)
var list = []int{}
var wg = &sync.WaitGroup{}

func main() {
	wg.Add(3)

	go func() {
		defer wg.Done()
		for _, v := range list {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()
		for v := range ch1 {
			ch2 <- v
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
