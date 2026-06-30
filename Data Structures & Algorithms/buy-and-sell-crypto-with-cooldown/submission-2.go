func maxProfit(prices []int) int {
    n := len(prices)
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, 2) // one for buy, one for not buy
	}
	for i := n-1; i >= 0; i-- {
		for allow2Buy := 1; allow2Buy >= 0; allow2Buy -- {
			if allow2Buy == 1 {
				buy := dp[i+1][0] - prices[i]
				cooldown := dp[i+1][1]
				dp[i][1] = max(buy, cooldown)
			} else {
				sell := prices[i]
				if i+2 < n {
					sell += dp[i+2][1]
				}
				cooldown := dp[i+1][0]
				dp[i][0] = max(sell, cooldown)
			}
		}
	}
	return dp[0][1]
}
