func findTargetSumWays(nums []int, target int) int {
    // bottom up method
	// dp = []map[int]int, dp[i][j] is using first i numbers, sum is j, have value ways
	// so dp[i+1][j+coins[i]] += dp[i][j]
	// dp[i+1][j-coins[i]] += dp[i][j]
	dp := make([]map[int]int, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make(map[int]int)
		//dp[i][0] = 1
	}
	dp[0][0] = 1 // this is the start point, from this, add next, or sub next, get two
	// and then get potential four. so we should only set dp[i][0] = 1
	for i := 0; i < len(nums); i++ {
		for canSum, count := range dp[i] {
			dp[i+1][canSum+nums[i]] += count
			dp[i+1][canSum-nums[i]] += count
		}
	}
	return dp[len(nums)][target]

}
