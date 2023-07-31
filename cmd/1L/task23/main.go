// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	s := []int{0, 5, 1, 8, 3, 2, 7, 6, 4, 7, 9, 3, 0, 2, 9}
	fmt.Println(s)
	fmt.Println(remove(s, 0))
}

func remove(s []int, i int) []int {
	if i < 0 || i >= (len(s)-1) {
		return s
	}

	s = append(s[:i], s[i+1:]...)
	return s
}
