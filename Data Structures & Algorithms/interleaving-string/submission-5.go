func isInterleave(s1 string, s2 string, s3 string) bool {
    // bottom up
	if len(s1) + len(s2) != len(s3) {return false}
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
        dp[i] = make([]bool, len(s2)+1)
    }
    dp[len(s1)][len(s2)] = true

	for i := len(s1); i >= 0; i-- {
		for j := len(s2); j >= 0; j-- {
			if i+1 <= len(s1) && s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
			}
			if j+1 <= len(s2) && s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}
	return dp[0][0]
}
