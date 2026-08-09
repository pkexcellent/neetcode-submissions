func search(nums []int, target int) bool {
	n := len(nums)
	l, r, m := 0, n-1, 0
	for l <= r {
		m = l + (r-l)/2
		//fmt.Println(l, r, m)
		if nums[m] == target {
			return true
		} 
		if nums[m] > nums[l] {
			// left part is monotonic
			if target < nums[m] && nums[l] <= target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[r] {
			// right part is monotonic
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			for l <= r && nums[l] == nums[m] {
				l++
			} 
			for l <= r && nums[r] == nums[m] {
				r--
			}
		}
	}
	return false
}
