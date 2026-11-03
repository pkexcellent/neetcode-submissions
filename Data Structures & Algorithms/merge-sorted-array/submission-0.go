func merge(nums1 []int, m int, nums2 []int, n int) {
	last := m+n-1
	pos1, pos2 := m-1, n-1
	for last >= 0 {
		if pos2 < 0 {
			break
		}
		if pos1 < 0 {
			for pos2 >= 0 {
				nums1[last] = nums2[pos2]
				last--
				pos2--
			}
			break
		}
		if nums1[pos1] > nums2[pos2] {
			nums1[last] = nums1[pos1]
			pos1--
		} else {
			nums1[last] = nums2[pos2]
			pos2--
		}
		last--
	}
}
