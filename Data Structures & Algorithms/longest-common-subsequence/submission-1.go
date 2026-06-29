func longestCommonSubsequence(text1 string, text2 string) int {
    // bottom up
	// dp[i][j] = x, i is text1's idx, j is text2', x is the common string length
	// dp[i][j] =  if text1[i] == text[j], dp[i][j] = dp[i-1][j-1] + 1
	// if not euqal, dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	m, n := len(text1)+1, len(text2)+1
	dp := make([][]int, m)
	
	for i, _ :=  range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])		
			}
		}
	}
	return dp[m-1][n-1]

}
