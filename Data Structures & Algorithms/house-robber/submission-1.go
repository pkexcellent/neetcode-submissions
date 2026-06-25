func rob(nums []int) int {
    // dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	// rob the i-1 one, or rob the i-2 + cur one
	n := len(nums)
	if n == 0 {return 0}
	if n == 1 {return nums[0]}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	}
	return dp[n-1]
}
