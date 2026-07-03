func findTargetSumWays(nums []int, target int) int {
    // use dfs, each number can be used as minus or positive
	// add memo
	// memo[i][j], i is the idx of the nums, j is the target, value is can reach ways
	memo := make(map[[2]int]int)

	var dfs func(sum int, idx int) int
	dfs = func(sum, idx int) int {
		if idx > len(nums) {
			return 0
		}
		if idx == len(nums) {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}
		if v, exist := memo[[2]int{idx, sum}]; exist {
			return v
		}

		// use this number as positive, // or use this as negetive
		rs := dfs(sum+nums[idx], idx+1) + dfs(sum-nums[idx], idx+1)
		
		memo[[2]int{idx, sum}] = rs
		return rs
	}
	return dfs(0, 0)
}

