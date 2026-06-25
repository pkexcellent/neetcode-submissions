func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
		// look at the video, it's clear
		if len(nums1) == 0 && len(nums2) == 0 {
			return 0.0
		}
		var short, long []int
		if len(nums1) > len(nums2) {
			short = nums2
			long = nums1
		} else {
			short = nums1
			long = nums2
		}
		var l, r = 0, len(short) // this means can pick [0, len(short)] elements from
		var totalLength = len(nums1) + len(nums2)
		for l <= r {
			mid1 := l + (r-l)/2 // mid1 is not the index, it means pick mid1 elements from short
			// if we picked i elements from short, we should pick ceil(totalLength/2) - i in another
			// mid2 means pick mid2 elements from long, it's not index
			mid2 := int(math.Ceil(float64(totalLength)/2)) - mid1

			// so picking mid1 elemenets, means shortL is element with index mid-1, shortR is with index mid
			shortL := math.MinInt64
			if mid1-1 >= 0 {
				shortL = short[mid1-1]
			}
			shortR := math.MaxInt64
			if mid1 <= len(short) - 1 {
				shortR = short[mid1]
			}
			longL := math.MinInt64 
			if mid2-1 >= 0 {longL = long[mid2-1]}
			longR := math.MaxInt64
			if mid2 <= len(long) - 1 {
				longR = long[mid2]
			}
			if shortL <= longR && longL <= shortR {
				if totalLength%2>0 {
					return float64(max(longL, shortL))
				} else {
					return float64(max(shortL, longL) + min(longR, shortR))/2.0
				}
			} else if shortL > longR {
				r = mid1-1
			} else {
				l = mid1+1
			}
		}
		return float64(-1)
}
