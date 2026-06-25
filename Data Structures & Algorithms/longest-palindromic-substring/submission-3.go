func longestPalindrome(s string) string {
    // DP method, not the best method, just for practicing
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}

	maxL := 0
	maxS := ""
	for r := 0; r < n; r++ {
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > maxL {
					maxL = r-l+1
					maxS = s[l:r+1]
				}
			}
		}
	}
	return maxS
}
