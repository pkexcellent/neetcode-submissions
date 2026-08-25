func canPartitionKSubsets(nums []int, k int) bool {
	// dfs
	total := 0
	largest := -1
	for _, num := range nums {
		total += num
		largest = max(largest, num)
	}
	if total%k != 0 {
		return false
	}
	each := total/k
	if largest > each {
		return false
	}
	groups := make([]int, k)
	sort.Slice(nums, func(i, j int) bool {return nums[i] > nums[j]})

	var dfs func(idx int) bool 
	dfs = func(idx int) bool {
		if idx == len(nums) {
			for _, v := range groups {
				if v != each {
					return false
				}
			}
			return true
		}
		for i := 0; i < k; i++ {
			if groups[i] + nums[idx] > each {
				continue
			}
			if i > 0 && groups[i] == groups[i-1] {
				continue
			}
			groups[i] += nums[idx]
			if dfs(idx+1) {
				return true
			}
			groups[i] -= nums[idx]
		}
		return false
	}
	return dfs(0)
}
