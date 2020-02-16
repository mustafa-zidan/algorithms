package sort

func InsertionSort(arr []int) {
    l := len(arr) - 1
    for i := l; i >= 0; i-- {
        for j := 0; j < i; j++ {
            if arr[j] > arr[i] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
}
