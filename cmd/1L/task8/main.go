// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

func main() {

}

func setByt(n int64, i uint, bitValue bool) int64 {
	if bitValue {
		n = n | (1 << i)
	} else {
		n = n & (1 << i)
	}
	return n
}
