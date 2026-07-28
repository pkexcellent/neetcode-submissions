func rotate(nums []int, k int) {
	k = k % len(nums)
	breakpoint := len(nums) - k
	rightP := nums[0:breakpoint]
	leftP := nums[breakpoint:len(nums)]
	x := append(leftP, rightP...)
	//nums = x // this doesn't work because golang slice in func is value copy
	// instead of reference copy
	copy(nums, x)
}
