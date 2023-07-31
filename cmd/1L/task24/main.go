// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func New(x, y int) *Point {
	return &Point{x, y}
}

func Distance(p, q *Point) float64 {
	return math.Sqrt(math.Pow(float64(q.x-p.x), 2) + math.Pow(float64(q.y-p.y), 2))
}

func main() {
	var p = New(1, 0)
	var q = New(2, 5)
	fmt.Println(Distance(p, q))
}
