func numSquares(n int) int {
	//dp[i] is the min number of perfect squares
	// for every j, dp[i] = min(dp[i], 1 + dp[i-j*j])
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i // all of 1s
	}
	for i := 0; i <= n; i++ {
		for k := 0; k*k <= i; k++ {
			dp[i] = min(dp[i], 1 + dp[i-k*k])
		}
	}
	return dp[n]
}
