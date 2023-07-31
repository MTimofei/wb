// Реализовать собственную функцию sleep.

package main

import "time"

func sleep(seconds time.Duration) {
	<-time.After(seconds * time.Second)
}
func main() {
	sleep(5)
}
