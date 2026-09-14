func minPathSum(grid [][]int) int {
	// dp
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j== 0 {
				continue
			} else if i == 0 {
				dp[j] = grid[i][j] + dp[j-1]
			} else if j == 0 {
				dp[j] = grid[i][j] + dp[j]
			} else {
				dp[j] = grid[i][j] + min(dp[j], dp[j-1])
			} 
		}
	}
	return dp[n-1]
}
