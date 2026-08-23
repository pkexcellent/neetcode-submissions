func canReach(s string, minJump int, maxJump int) bool {
	// dp
	n := len(s)
	if s[n-1] != '0' {return false}

	dp := make([]bool, n)
	dp[0] = true
	for i := 0; i < n; i++ {
		if dp[i] != true {continue}
		for j := i + minJump; j <= i + maxJump && j < n; j++ {
			if s[j] == '0' {
				dp[j] = true
			}
			if j == n-1 && dp[j] {
				return true
			}
		}
	}
	return dp[n-1]
}
