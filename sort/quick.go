package main

import (
	"fmt"
)

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivot := doPivot(arr, low, high)
	fmt.Println(arr[low:high])
	quickSort(arr, low, pivot-1)
	quickSort(arr, pivot+1, high)
}

func doPivot(arr []int, low, high int) int {
	leftwall := low
	pivot := arr[high]
	for leftwall < high {
		if pivot > arr[leftwall] {
			fmt.Println("swapping", arr[leftwall], arr[low])
			arr[low], arr[leftwall] = arr[leftwall], arr[low]
			low++
		}
		leftwall++
	}
	arr[high], arr[low] = arr[low], arr[high]
	return low
}

func QuickSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	quickSort(arr, 0, length-1)
}

func main() {
	arr := []int{11, 16, 2, 8, 1, 9, 4, 7}
	fmt.Println(arr)
	QuickSort(arr)
	fmt.Println(arr)
}
