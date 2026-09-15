func stoneGame(piles []int) bool {
	// use a dp method
	//  dp[i][j] to store the max alice can get within i to j
	//  then dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
	// initial status, dp[i][i] is piles[i]
	n := len(piles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for width := 2; width <= n; width++ {
		for i := 0; i + width -1 < n; i++ {
			j := i + width - 1
			dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
		}
	}
	return dp[0][n-1] > 0 
}
