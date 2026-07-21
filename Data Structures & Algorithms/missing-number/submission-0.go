func missingNumber(nums []int) int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		diff += i - nums[i]
	}
	diff += len(nums)
	return diff
}
