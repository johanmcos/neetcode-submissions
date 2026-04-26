func canPartition(nums []int) bool {
    var totalSum int
	for _, num := range nums {
		totalSum += num
	}

	// an odd sum can't be evenly divided
	if totalSum % 2 != 0 {
		return false
	}
	target := totalSum / 2

	sumMap := make(map[int]struct{}, len(nums))

	for _, num := range nums {
    	newSums := make(map[int]struct{})
		for val := range sumMap {
			newSums[val+num] = struct{}{}
		}
		for val := range newSums {
			sumMap[val] = struct{}{}
		}
		sumMap[num] = struct{}{}
		if _, exists := sumMap[target]; exists {
			return true
		}
	}
	return false
}
