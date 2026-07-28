func rotate(nums []int, k int) {
	// three times revert
	n := len(nums)
	k = k%n
	revert(nums, 0, n-1)
	revert(nums, 0, k-1)
	revert(nums, k, n-1)
}

func revert(nums []int, l, r int) {
	for i, j := l, r; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
