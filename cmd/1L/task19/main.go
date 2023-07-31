// Разработать программу,
// которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

package main

import "fmt"

func main() {
	var word string = "главрыба"

	fmt.Println(
		reverse1(word),
		word,
	)
	fmt.Println(
		reverse2(word),
		word,
	)

}

func reverse1(s string) string {
	run := []rune(s)
	result := []rune("")
	for i := len(run) - 1; i >= 0; i-- {
		result = append(result, run[i])
	}
	return string(result)
}

func reverse2(s string) string {
	runs := []rune(s)
	length := len(runs)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runs[i], runs[j] = runs[j], runs[i]
	}
	return string(runs)

}
