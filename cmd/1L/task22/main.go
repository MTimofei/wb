// Разработать программу, которая перемножает,
// делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите первое число")
	i := readNumber()
	fmt.Println("Введите второе число")
	n := readNumber()
	fmt.Println("Введите знак операцию")
	sig := readString()
	fmt.Println(Calculated(i, n, sig))
}

func Calculated(i, n float64, sig string) float64 {

	switch {
	case sig == "+":
		return i + n
	case sig == "-":
		return i - n
	case sig == "*":
		return i * n
	case sig == "/":
		return i / n
	}
	return 0
}

func readString() string {
	var s string
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		fmt.Println("Ошибка чтения строки:", err)
		return ""
	}
	return s
}

func readNumber() float64 {
	var n float64
	_, err := fmt.Scanf("%f", &n)
	if err != nil {
		fmt.Println("Ошибка чтения числа:", err)
		return 0
	}
	return n
}
