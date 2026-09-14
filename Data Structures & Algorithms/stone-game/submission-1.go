func stoneGame(piles []int) bool {
	// dfs way, either pick from first or last
	// 
	memo := make(map[[2]int]int)
	var dfs func(l, r int, alice bool) int
	dfs = func(l, r int, alice bool) int {
		if l == r {
			if alice {
				return piles[l]
			} else {
				return -piles[l]
			}
		}
		if v, exist := memo[[2]int{l, r}]; exist {
			return v
		}
		diff := 0
		if alice {
			// try to make the different big enough
			diff = max(piles[l] - dfs(l+1, r, false), 
				piles[r] - dfs(l, r-1, false))
		} else {
			// try to make the diff small enough
			diff = min(piles[l] - dfs(l+1, r, true), 
				piles[r] - dfs(l, r-1, true))
		}
		memo[[2]int{l, r}] = diff
		return diff
	}
	rs := dfs(0, len(piles)-1, true)
	return rs > 0
}
