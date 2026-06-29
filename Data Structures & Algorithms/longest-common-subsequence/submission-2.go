func longestCommonSubsequence(text1 string, text2 string) int {
    // explore a top down way, using dfs
	memo := make([][]int, len(text1))
	for i, _ := range memo {
		memo[i] = make([]int, len(text2))
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		rs := 0
		if text1[i] == text2[j] {
			rs = 1 + dfs(i+1, j+1)
		} else {
			rs = max(dfs(i+1, j), dfs(i, j+1))
		}
		memo[i][j] = rs
		return rs
	}
	return dfs(0, 0)
}
