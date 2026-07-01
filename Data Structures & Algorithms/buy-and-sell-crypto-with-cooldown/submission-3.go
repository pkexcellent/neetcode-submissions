func maxProfit(prices []int) int {
    n := len(prices)
    dp1_buy, dp1_sell := 0, 0
    dp2_buy := 0

    for i := n - 1; i >= 0; i-- {
        dp_buy := max(dp1_sell - prices[i], dp1_buy)
        dp_sell := max(dp2_buy + prices[i], dp1_sell)
        dp2_buy, dp1_sell = dp1_buy, dp_sell
        dp1_buy = dp_buy
    }

    return dp1_buy
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
