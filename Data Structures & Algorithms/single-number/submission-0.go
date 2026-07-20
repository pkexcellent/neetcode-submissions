func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	rs := nums[0]
	for i := 1; i < len(nums); i++ {
		rs = rs^nums[i]
	}
	return rs
}
