package sort

func InsertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        for j := i; j > 0; j-- {
            if arr[j-1] > arr[j] {
                arr[j-1], arr[j] = arr[j], arr[j-1]
            } else {
                // if there is no swapping then stop this cycle since rest of
                // elements are sorted
                break
            }
        }
    }
}
