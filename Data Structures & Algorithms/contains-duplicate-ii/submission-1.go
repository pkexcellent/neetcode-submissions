func containsNearbyDuplicate(nums []int, k int) bool {
	// not to degenerate this algorithm
	// becasue if we have a very long nums list
	// the solution1 with map has to be very huge
	// and probable, every ele has to be calculated
	m := make(map[int]struct{})
	l := 0
	for r := 0; r < len(nums); r++ {
		if r - l > k {
			delete(m, nums[l])
			l++
		}
		if _, exist := m[nums[r]]; exist {
			return true
		}
		m[nums[r]] = struct{}{}
	}
	return false
}
