func change(amount int, coins []int) int {
    sort.Ints(coins)
	// use DP
	// dp[cointIdx][sum] = dp[coinIdx+1][sum] + dp[coinIdx][sum-coins[idx]]
	// initial status, dp[*][0] = 1; dp[len(coins)][amount] = 0 
	dp := make([][]int, len(coins)+1)
	for i, _ := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := len(coins)-1; i >= 0; i-- {
		for sum := 1; sum <= amount; sum++ {
			if sum >= coins[i] {
				dp[i][sum] = dp[i+1][sum] + dp[i][sum-coins[i]]
			} else {
				dp[i][sum] = dp[i+1][sum]
			}	
		}
	}
	return dp[0][amount]  
}