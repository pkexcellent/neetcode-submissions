func numDecodings(s string) int {
    // DP bottom up
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	if s[len(s)-1] == '0' {
		dp[len(s)-1] = 0
	} else {
		dp[len(s)-1] = 1
	}
	for i := len(s)-2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] += dp[i+1]
			if s[i] == '1' || (s[i] == '2' && s[i+1] < '7') {
				dp[i] += dp[i+2]
			}
		}

	}
	return dp[0]
}
