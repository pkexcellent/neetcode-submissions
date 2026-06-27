func lengthOfLIS(nums []int) int {
    // top-down
	// pick i dfs(i+1, largest), or not pick i, then dfs(i+1, largest)
	// memo, has to 2-dimensions 
	memo := make(map[[2]int]int)
	var dfs func(int, int) int
	dfs = func(idx, largestSoFar int) int {
		if idx == len(nums) {
			return 0
		}
		if v, exist := memo[[2]int{idx, largestSoFar}]; exist {
			return v
		}
		// pick
		var pickCount, notPickCount int
		if nums[idx] > largestSoFar {
			pickCount = 1 + dfs(idx+1, nums[idx])
		}
		// not pick
		notPickCount = dfs(idx+1, largestSoFar)
		memo[[2]int{idx, largestSoFar}] = max(pickCount, notPickCount)
		return memo[[2]int{idx, largestSoFar}]
	}
	rs := dfs(0, math.MinInt)
	//fmt.Println(memo)
	return rs
}
