// Поменять местами два числа без создания временной переменной

package main

type space struct{}

func main() {
	var n space
	var i space
	n, i = i, n
}
