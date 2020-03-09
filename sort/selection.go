package sort

// O(n^2)
// (n-1) + (n-2) + (n-3) +... + 1  = n(n-1)/2
func SelectionSort(arr []int) {
    for i, v := range arr {
        smallestIndex, smallestValue := i, v
        for j := i; j < len(arr); j++ {
            if smallestValue > arr[j] {
                smallestIndex, smallestValue = j, arr[j]
            }
        }
        arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
    }
}
