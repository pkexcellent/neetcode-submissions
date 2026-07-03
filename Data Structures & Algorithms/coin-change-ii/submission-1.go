func change(amount int, coins []int) int {
    // use dfs, pick the coin or not, repeatly

	// really tracking the combinations will NZEC, comb := [][]int{}
	// so just count
	// brute force dfs will TLE, so we use a memo to record the history
	memo := make([][]int, len(coins))
	for i, _ := range memo {
		memo[i] = make([]int, amount+1)
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(idx int, remain int) int
	dfs = func(idx int, remain int) int {
		if idx >= len(coins) {return 0}
		if remain == 0 {
			//comb = append(comb, pick)
			return 1
		}
		if remain < 0 {
			return 0
		}
		if memo[idx][remain] != -1 {
			return memo[idx][remain]
		}
		res := 0
		// pick the coin at current idx
		res += dfs(idx, remain-coins[idx])
		// pick next one
		res += dfs(idx+1, remain)
		memo[idx][remain] = res
		return res

	}
	
	x := dfs(0, amount)
	//fmt.Println(memo)
	return x
}
