package sort

func quickSort(arr []int, low, high int) {
    if low >= high {
        return
    }
    pivot := doPivot(arr, low, high)
    quickSort(arr, low, pivot-1)
    quickSort(arr, pivot+1, high)
}

func doPivot(arr []int, low, high int) int {
    leftwall := low
    pivot := arr[high]
    for leftwall < high {
        if pivot > arr[leftwall] {
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
