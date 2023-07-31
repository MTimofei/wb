// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"flag"
	"fmt"
	"time"
)

var workHours = flag.Int64("WorkHours", 5, "application working hours")

func main() {
	var ch = make(chan uint64)
	var ok bool = true

	go func() {
		time.Sleep(time.Duration(*workHours) * time.Second)
		ok = false
	}()

	go func() {
		var n uint64
		for ok {
			ch <- n
			n++
		}
	}()

	go func() {
		for ok {
			select {
			case n := <-ch:
				fmt.Println(n)
			default:
			}
		}
	}()

	for ok {
	}
}
