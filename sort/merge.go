package sort

func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    pivot := (len(arr)) / 2
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

func MergeInPlaceSort(arr []int) {

}
