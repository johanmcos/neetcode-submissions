import "slices"

func getConcatenation(nums []int) []int {
    return slices.Concat(nums, nums)
}
