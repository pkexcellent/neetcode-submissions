func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if latestIdx, exist := m[num]; exist {
			if i - latestIdx <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}
