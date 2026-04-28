func twoSum(nums []int, target int) []int {
	i, j := 0, 1
	for nums[i] + nums[j] != target {
		if j < len(nums) - 1 {
			j++
		} else {
			i++
			j = i + 1
		}
	}
    return []int{i, j}
}
