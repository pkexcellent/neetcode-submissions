func maxSlidingWindow(nums []int, k int) []int {
    if k > len(nums) {
		return nil
	}
	var maxInWindow = make([]int, len(nums)-k+1)
	var l, r = 0, k-1
	for r < len(nums) {
		maxInWindow[l] = nums[l]
		for i := l; i<=r; i++ {
			if nums[i] > maxInWindow[l] {
				maxInWindow[l] = nums[i]
			}
		}
		l++
		r++
	}
	return maxInWindow
}
