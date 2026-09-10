func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for sum := 1; sum <= target; sum++ {
		for _, num := range nums {
			if sum - num >= 0 {
				dp[sum] += dp[sum-num]
			}
		}
	}
	return dp[target]
}
