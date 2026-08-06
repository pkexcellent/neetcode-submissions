func splitArray(nums []int, k int) int {
	// get sum of the nums, and get the average sub array
	// try to make every array closest to the even
	sum := 0
	n := len(nums)
	for _, num := range nums {
		sum += num
	}
	even := (sum + k - 1)/k
	memo := make(map[[2]int]int)
	var dfs func(idx int, remainK int) int
	dfs = func(idx int, remainK int) int {
		if (remainK == 0 && idx != n) || (remainK > 0 && idx == n) {
			return 1000000000 // invalid case
		}
		if v, exist := memo[[2]int{idx, remainK}]; exist {
			return v
		}
		subsum := 0
		rs := 10000000
		for i := idx; i <= n-remainK; i++ {
			subsum += nums[i]
			if i == n-1 {	
				return subsum
			} else if subsum >= even || i == n-remainK {
				//maxSubsum := 0
				restMaxSumWith := dfs(i+1, remainK-1)
				restMaxSumWith = max(subsum, restMaxSumWith)
				restMaxSumWithout := 1000000000
				if i > idx {
					restMaxSumWithout = dfs(i, remainK-1)
					restMaxSumWithout = max(subsum-nums[i], restMaxSumWithout)
				}
				rs = min(rs, min(restMaxSumWithout, restMaxSumWith))
			} 
		}
		memo[[2]int{idx, remainK}] = rs
		return rs 
	}

	maxsum := dfs(0, k)
	return maxsum
}
