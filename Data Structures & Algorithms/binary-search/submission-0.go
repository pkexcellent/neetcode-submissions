func search(nums []int, target int) int {
	var l, r = 0, len(nums)-1
	for l <= r {
		if nums[l + (r-l)/2] == target {
			return l + (r-l)/2
		} else if nums[l + (r-l)/2] > target {
			r = l + (r-l)/2 -1
		} else {
			l = l + (r-l)/2 +1
		}
	}
	return -1
}
