func findKthLargest(nums []int, k int) int {
	//sort.Slice(nums, func(i, j int) bool {return nums[i] > nums[j]})
	//return nums[k-1]
	k = len(nums)-k
	var quickSelect func(l, r, k int) int
	quickSelect = func(l, r, k int) int {
		pivot := nums[r]
		p := l
		for i := l; i < r; i++ {
			if nums[i] <= pivot {
				nums[i], nums[p] = nums[p], nums[i]
				p++
			}
		}
		nums[p], nums[r] = nums[r], nums[p]
		if p == k {
			return nums[p]
		} else if p < k {
			return quickSelect(p+1, r, k)
		} else {
			return quickSelect(l, p-1, k)
		}
	}
	return quickSelect(0, len(nums)-1, k)
}
