// Write a program that competitively
// calculates the value of squares of numbers taken from
// array (2,4,6,8,10) and outputs their squares to stdout.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch = make(chan int)
	var wg sync.WaitGroup
	var list = []int{2, 4, 6, 8, 10}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, n := range list {
			ch <- n
		}
	}()

	for i := 0; i < len(list); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := <-ch
			fmt.Println(n * n)
		}()
	}
	wg.Wait()
}
