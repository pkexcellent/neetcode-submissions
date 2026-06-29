func longestCommonSubsequence(text1 string, text2 string) int {
    // bottom up
	// dp[i][j] = x, i is text1's idx, j is text2', x is the common string length
	// dp[i][j] =  if text1[i] == text[j], dp[i][j] = dp[i-1][j-1] + 1
	// if not euqal, dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	
	for i, _ :=  range dp {
		dp[i] = make([]int, n)
	}

	if text1[0] == text2[0] {dp[0][0] = 1}
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				if i-1>=0 && j-1>=0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				if i-1>=0 && j-1>=0 {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				} else if i-1>=0 {
					dp[i][j] = dp[i-1][j]
				} else if j-1>=0{
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = 0
				}
				
			}
		}
	}
	for i := 0; i < m; i++ {fmt.Println(dp[i])}
	return dp[m-1][n-1]

}
