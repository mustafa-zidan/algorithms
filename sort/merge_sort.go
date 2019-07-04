package main

import (
	"fmt"
)

func MergeSort(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		return arr
	}
	pivot := (length) / 2
	return merge(MergeSort(arr[:pivot]), MergeSort(arr[pivot:]))
}

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	result := make([]int, size, size)
	for k, _ := range result {
		if i == len(left) {
			result[k] = right[j]
			j++
		} else if j == len(right) {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}
	return result
}

func main() {
	arr := []int{11, 16, 2, 8, 1, 9, 4, 7}
	fmt.Println(arr)
	fmt.Println(MergeSort(arr))
}
