func minCostClimbingStairs(cost []int) int {
    // should be DP
	// the cost of ith is the min of (i-1)th +cost[i-1] or (i-2)th + cost
	n := len(cost)
	if n == 0 {return 0}
	if n == 1 {return cost[0]}

	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
	}
	return dp[len(cost)]
}
