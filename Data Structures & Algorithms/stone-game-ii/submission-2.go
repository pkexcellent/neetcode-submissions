func stoneGameII(piles []int) int {
	// use top-down dp
	// use dfs method to illstrate the game flow
	// while picking, try maximun Alice's and minimun Bob'
	n := len(piles)
	memo := make([][][2]int, n)
	for i, _ := range memo {
		memo[i] = make([][2]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = [2]int{math.MinInt32, math.MinInt32}
		}
	}
	var dfs func(idx int, M int, isAlice bool) int
	dfs = func(idx int, M int, isAlice bool) int {
		if idx == n {
			return 0
		}
		memoIdxAlice := 1
		if !isAlice {
			memoIdxAlice = 0
		}
		if memo[idx][M][memoIdxAlice] != math.MinInt32 {
			return memo[idx][M][memoIdxAlice]
		}
		final := math.MinInt32
		if isAlice {
			aliceGet := 0
			for i := idx; i < min(idx + 2*M, n); i++ {
				aliceGet += piles[i]
				final = max(final, aliceGet + dfs(i+1, max(i-idx+1, M), false))
				
			}
		} else {
			final = math.MaxInt32
			for i := idx; i < min(idx + 2*M, n); i++ {
				final = min(final, dfs(i+1, max(i-idx+1, M), true))
			}
		}
		memo[idx][M][memoIdxAlice] = final
		return final
	}
	return dfs(0, 1, true)
}
