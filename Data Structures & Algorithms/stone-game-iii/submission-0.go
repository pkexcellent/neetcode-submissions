func stoneGameIII(stoneValue []int) string {
	// pick stone from rear to head
	//dp[i] is the (max value that Alice can get) - (max bob can get)
	// update dp[i] iterately, 
	// for j start from i to t+3, update dp[i] as the (max sum so far) - (dp[j+1])
	n := len(stoneValue)
	dp := make([]int, n+1)
	dp[n] = 0
	for i := n-1; i >= 0; i-- {
		aliceSum := 0 
		dp[i] = math.MinInt32
		for j := i; j < min(n, i+3); j++ {
			aliceSum += stoneValue[j]
			dp[i] = max(dp[i], aliceSum - dp[j+1])
		}
	}
	if dp[0] > 0 {
		return "Alice"
	} else if dp[0] < 0 {
		return "Bob"
	} else {
		return "Tie"
	}
}
