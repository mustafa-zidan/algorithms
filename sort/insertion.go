package main

import (
	"fmt"
)

func InsertionSort(arr []int) {
	l := len(arr)-1
	for i := l; i >= 0  ; i-- {
		for j:= 0 ; j < i ; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}


func main() {
	arr := []int{11, 16, 2, 8, 1, 9, 4, 7}
	fmt.Println(arr)
	InsertionSort(arr)
	fmt.Println(arr)
}
