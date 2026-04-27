func findMaxConsecutiveOnes(nums []int) int {
	var maxConsecutive, currentConsecutive int
    for _, num := range nums {
        if num == 1 {
            currentConsecutive++
            maxConsecutive = max(maxConsecutive, currentConsecutive)
        } else {
            currentConsecutive = 0
        }
    }

    return maxConsecutive
}
