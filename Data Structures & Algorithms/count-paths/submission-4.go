func uniquePaths(m int, n int) int {
    // use rolling DP
	dp := make([]int, n)
	for c := 0; c < n; c++ {
		dp[c] = 1
	}
	//fmt.Println(dp)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j] // dp[j] is the last round, dp[j-1] is the last round
		}
		//fmt.Println(dp)
	}
	return dp[n-1]
}
