func searchInsert(nums []int, target int) int {
	// binary search
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	// -1,0,2,4,6,8  | 5
	// l, r, m
	// 0, 5, 2
	// 3, 5, 4
	// 3, 4, 3
	// 4, 4, 4
	// 4, 3
	return r + 1
}
