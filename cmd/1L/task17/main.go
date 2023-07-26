// Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 6

	index := binarySearch(arr, target)
	fmt.Println(index)
}

// func binarySearch(arr []int, target int) int {
// 	i := len(arr) / 2
// 	switch {
// 	case arr[i] > target:
// 		return i - binarySearch(arr[:i+1], target)
// 	case arr[i] < target:
// 		return i + binarySearch(arr[i:], target)
// 	case arr[i] == target:
// 		return i
// 	default:
// 		return -1
// 	}
// }

func binarySearch(arr []int, target int) int {
	index := sort.SearchInts(arr, target)
	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1
}
