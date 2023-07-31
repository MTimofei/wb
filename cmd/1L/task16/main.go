// Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

type By []int

func (a By) Len() int           { return len(a) }
func (a By) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a By) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	arr := []int{0, 5, 1, 8, 3, 2, 7, 6, 4, 7, 9, 3, 0, 2, 9}
	fmt.Println(arr)

	sort.Sort(By(arr))
	fmt.Println(arr)

	arr = []int{0, 5, 1, 8, 3, 2, 7, 6, 4, 7, 9, 3, 0, 2, 9}
	fmt.Println(arr)

	quicksort(arr)
	fmt.Println(arr)

}

func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)/2]
	left := 0
	right := len(arr) - 1

	for left <= right {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	quicksort(arr[:right+1])
	quicksort(arr[left:])
}
