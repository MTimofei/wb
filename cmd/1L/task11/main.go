// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

type space struct{}

func main() {
	set1 := []int{10, 7, 3, 9, 5}
	set2 := []int{4, 5, 9, 7, 8}

	fmt.Println(intersection(set1, set2))
}

func intersection(nums1, nums2 []int) []int {
	m := make(map[int]space)
	nums := make([]int, 0, 0)
	for _, n := range nums1 {
		m[n] = space{}
	}

	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			nums = append(nums, n)
		}
	}

	return nums
}
