package search

func BuyLots(lots []int, target int) []int {
    i, j, sum := 0, 1, 0
    for j < len(lots) {
        if sum += lots[j]; sum == target {
            return lots[i:j]
        } else if sum > target {
            sum -= lots[i]
            i++
        } else if sum < target {
            j++
        }
    }
    return nil
}
