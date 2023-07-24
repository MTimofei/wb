// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import (
	first "github.com/wb/cmd/1L/task1/1/action"
	second "github.com/wb/cmd/1L/task1/1/action"
)

func main() {

	// пример первого варианта
	var a2 first.Action

	a2.Age()
	a2.Walk()

	//пример второго варианта
	var a second.Action

	a.Age()
	a.Walk()
}
