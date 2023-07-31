// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(refersi("snow dog sun"))
}

func refersi(s string) string {
	sList := strings.Split(s, " ")
	l := len(sList)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		sList[i], sList[j] = sList[j], sList[i]
	}
	return strings.Join(sList, " ")
}
