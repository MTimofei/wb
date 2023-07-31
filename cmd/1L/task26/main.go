// Разработать программу, которая проверяет,
// что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(isUniqueChars(str1), isUniqueChars(str2), isUniqueChars(str3))
}

func isUniqueChars(str string) bool {
	chars := strings.Split(str, "")
	m := make(map[string]struct{})
	for _, c := range chars {
		if _, ok := m[c]; ok {
			return false
		}
		m[c] = struct{}{}
	}
	return true
}
