import "slices"

func lengthOfLIS(nums []int) int {
    lens := make([]int, len(nums))
	for i, val := range nums {
		lens[i] = 1
		for j := i - 1; j >= 0; j-- {
			if val > nums[j] {
				lens[i] = max(lens[i], lens[j]+1)
			}
		}
	}

	return slices.Max(lens)
}
