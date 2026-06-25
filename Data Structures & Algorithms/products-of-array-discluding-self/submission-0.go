func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var ans = make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	var backward = 1
	for i := len(nums) - 2; i >= 0; i-- {
		backward = backward * nums[i+1]
		ans[i] = ans[i] * backward
	}
	return ans
}
