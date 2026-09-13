func integerBreak(n int) int {
	// dp[i] to store the max value
	// for every 0<k<i; dp[i] = max(dp[i], max(k, dp[k])*max(i-k, dp[i-k]))
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i-1 // set init number (1*(i-1))
	}
	for i := 2; i <= n; i++ {
		for k := 1; k < i; k++ {
			// when dp[k], k > 3, there is a choice that whether break it or not
			// this is not like dp[2], dp must be broken into 1 + 1
			// but when it is dp[4], when calculating (2, 2), we don't need to break it to
			// dp[2]*dp[2], we can use max(2, dp[2])*max(2, dp[2])
			// so it's max(k, dp[k]) * max(i-k, dp[i-k]) 
			dp[i] = max(dp[i], max(k, dp[k])*max(i-k, dp[i-k]))
		}
	}
	//fmt.Println(dp)
	return dp[n]
}
