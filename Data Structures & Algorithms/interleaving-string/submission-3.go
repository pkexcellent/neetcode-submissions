func isInterleave(s1 string, s2 string, s3 string) bool {
    // brute force dfs will exceed the time limit
	// so adding memo
	// need len+1 size, because in DFS, i and j and be len+1, we still need to record the status
	// of len+1, and valid j
	if len(s1)+len(s2) != len(s3) {return false}
	memo := make([][]int, len(s1)+1)
	for i, _ := range memo {
		memo[i] = make([]int, len(s2)+1)
		for j := 0; j < len(s2)+1; j++ {
			memo[i][j] = -1 // -1 means not set
		}
	}
	var dfs func(i, j, k int) bool 
	dfs = func(i, j, k int) bool {
		if k == len(s3) && i == len(s1) && j == len(s2) {
			return true
		}

		if memo[i][j] != -1 {
			return memo[i][j] == 2
		}
		if i < len(s1) && s1[i] == s3[k] {
			if dfs(i+1, j, k+1) {
				memo[i][j] = 2
				return true
			}
		}
		if j < len(s2) && s2[j] == s3[k] {
			if dfs(i, j+1, k+1) {
				memo[i][j] = 2
				return true
			}
		}
		memo[i][j] = 1
		return false
	}
	return dfs(0, 0, 0)
}
