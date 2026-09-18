func lastStoneWeightII(stones []int) int {
	// knapsack dp[i][t] i is idx, t is capacity weight, dp is the maximum sum of stones
	total := 0
	for _, stone := range stones {
		total += stone
	}
	n := len(stones)
	target := total/2
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 0
	for capacity := 0; capacity < target + 1; capacity++ {
		for i := 1; i <= n; i++ {
			if stones[i-1] <= capacity {
				dp[i][capacity] = max(dp[i-1][capacity], 
						dp[i-1][capacity-stones[i-1]] + stones[i-1])
			} else {
				dp[i][capacity] = dp[i-1][capacity]
			}
		}
	}
	maxWeight := dp[n][target]
	return abs(maxWeight - (total-maxWeight))

}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
