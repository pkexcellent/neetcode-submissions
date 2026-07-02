func maxProfit(prices []int) int {
    // A true bottom up method
	// 3 states are needed:
	// hold[i] is the max profit if hold stock at i day
	// sold[i] is the max profit is sell stock at i day
	// ableBuy[i] is the max profit if able to buy stock at i day
	// hold[i] = max(hold[i-1], ableBuy[i-1]-price[i])
	// sold[i] = hold[i-1] + price[i]
	// ableBuy[i] = max(sold[i-1], ableBuy[i-1])
	n := len(prices)
	hold := make([]int, n)
	sold := make([]int, n)
	ableBuy := make([]int, n)
	hold[0] = -prices[0]
	sold[0] = 0
	ableBuy[0] = 0
	for i := 1; i < n; i++ {
		hold[i] = max(hold[i-1], ableBuy[i-1] - prices[i])
		sold[i] = hold[i-1] + prices[i]
		ableBuy[i] = max(ableBuy[i-1], sold[i-1])
	}
	return max(hold[0], sold[n-1], ableBuy[n-1])
}
