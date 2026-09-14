func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp[i][j] is the method can reach to ij
	//  dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// if ij is 1, then dp[i][j] is 0; 
	dp := make([][]int, len(obstacleGrid))
	for i, _ := range dp {
		dp[i] = make([]int, len(obstacleGrid[i]))
	} 
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	dp[0][0] = 1
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if i > 0 {
					dp[i][j] += dp[i-1][j]
				} 
				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
