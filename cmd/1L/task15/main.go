// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

package main

import (
	"fmt"
	"io"
	"strings"
)

var justString string
var lr io.LimitedReader

// чтобы не читать все строку для копирования первых 100 быйт(тем самым не выделять лишнюю пямфть)

func someFunc() {
	buf := make([]byte, 100)
	if _, err := io.ReadFull(&lr, buf); err != nil {
		fmt.Println("Error reading string:", err)
		return
	}

	justString = string(buf)
}
func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func init() {
	lr = io.LimitedReader{R: strings.NewReader(createHugeString(1 << 10)), N: 100}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
