func replaceElements(arr []int) []int {
	for i:=0; i < len(arr) - 1; i++ {
		var maxRight int
		for j := i + 1; j < len(arr); j++ {
			maxRight = max(arr[j], maxRight)
		}
		arr[i] = maxRight
	}
	arr[len(arr)-1] = -1

	return arr
}
