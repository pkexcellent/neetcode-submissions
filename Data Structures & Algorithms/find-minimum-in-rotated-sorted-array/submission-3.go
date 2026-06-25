func findMin(nums []int) int {
	var ans = nums[0]
	var l, r = 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		//fmt.Println(l, mid, r)
		if nums[mid] >= nums[l] {
			// if left is monotonic
			ans = min(ans, nums[l])
			l = mid + 1
		} else { //if nums[mid] < nums[r] {
			// if right is monotonic
			ans = min(ans, nums[mid])
			r = mid - 1
		}
	}
	return ans
}
