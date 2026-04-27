func hasDuplicate(nums []int) bool {
	var numsMap = make(map[int]struct{})
	for _, num := range nums {
		_, exists := numsMap[num]
		if exists {
			return true
		}
		numsMap[num] = struct{}{}
	}
    return false
}
