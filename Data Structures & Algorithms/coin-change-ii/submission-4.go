func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
	dp[0] = 1

	// dp[amount] = dp[amount - coin] <- for each coin
	for i := len(coins)-1; i >= 0; i-- {
		for sum := 1; sum <= amount; sum++ {
			if sum >= coins[i] {
				dp[sum] = dp[sum] + dp[sum-coins[i]]
			}
		}
	}
	//fmt.Println(dp)
	return dp[amount]
}
