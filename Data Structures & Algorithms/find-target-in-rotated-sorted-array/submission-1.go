func search(nums []int, target int) int {

	var l, r = 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[l] {
			// it's monotonic on left
			if target > nums[mid]|| target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// it's monotonic on right
			if target > nums[r] || target < nums[mid]{
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
