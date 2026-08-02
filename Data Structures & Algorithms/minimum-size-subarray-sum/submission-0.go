func minSubArrayLen(target int, nums []int) int {
	// use a flex window, adjusting l and r
	l, r, n := 0, 0, len(nums)
	if n == 0 {
		return 0
	}
	sum := nums[0]
	window := n+1
	for l <= r && r < n {
		//fmt.Println(l, r, sum, window)
		if sum >= target {
			window = min(window, r-l+1)
			sum -= nums[l]
			l++
			if l > r {
				r++
			}
		} else if sum < target {
			r++
			if r < n {
				sum += nums[r]
			}
		}
	}
	if window == n+1 {
		return 0
	}
	return window
}
