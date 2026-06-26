func coinChange(coins []int, amount int) int {
    // use a DP method
	// dp[i] means using how many coins can fill amount i
	// then dp[j] = min(dp[j], 1 + dp[j-x]) for x in each coin domination
	dp := make([]int, amount+1)
	for i := 0; i< amount+1; i++ {
		dp[i] = amount+1
	}
	dp[0] = 0
	for i:=1; i<amount+1; i++ {
		for _, coin := range coins {
			if i - coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin] + 1)
			}
		}
	}
	if dp[amount]==amount+1 {
		return -1
	}
	return dp[amount]
}
