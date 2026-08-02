func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			l++
			nums[l] = nums[i]
		}
	}
	return l+1
}
