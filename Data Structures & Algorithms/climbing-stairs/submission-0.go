func climbStairs(n int) int {
    // classical DP
	// dp[i] is the methods reach to position i
	// so dp[i+1] = dp[i] + dp[i-1]
	if n == 0 {return 0}
	if n == 1 {return 1}
	dp := make([]int, n)
	dp[0], dp[1] = 1, 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
