func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([]int, len(obstacleGrid[0]))
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	dp[0] = 1
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				if j > 0 {
					dp[j] = dp[j] + dp[j-1]
				} 
			}
		}
	}
	return dp[len(obstacleGrid[0])-1]
}

