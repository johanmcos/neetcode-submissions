func removeElement(nums []int, val int) int {
    k := 0 // track the index for valid elements
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
