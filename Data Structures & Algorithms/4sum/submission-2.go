func fourSum(nums []int, target int) [][]int {
	// sort it, fix the 1st and 2nd, and use 2 pointers
	sort.Ints(nums)
	rs := [][]int{}
	n := len(nums)
	for i := 0; i < n-3; i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i+1; j < n-2; j++ {
			// should judge j > i+1, instead of j > 1, as the first j after i must be calculated once
			if j > i+1 && nums[j] == nums[j-1] { 
				continue
			}
			l, r := j+1, n-1
			for l < r {
				if nums[i] + nums[j] + nums[l] + nums[r] == target {
					rs = append(rs, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for l < n && nums[l] == nums[l-1] {
						l++
					}
					r--
					for r > j && nums[r] == nums[r+1] {
						r--
					}

				} else if nums[i] + nums[j] + nums[l] + nums[r] < target {
					l++
					for l < n && nums[l] == nums[l-1] {
						l++
					}
				} else {
					r--
					for r > j && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	} 
	return rs
}
