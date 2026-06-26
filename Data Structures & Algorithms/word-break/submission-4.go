func wordBreak(s string, wordDict []string) bool {
    // top-down
	// build a wordDict dict
	wd := make(map[string]bool)
	maxL := 0
	for _, word := range wordDict {
		wd[word] = true
		if len(word) > maxL {maxL = len(word)}
	}
	// memo, memo[i] indicate memo[i] is true
	memo := make(map[int]bool)

	// purely loop will time exceed, so add a len to skip uncessary loop

	var dfs func(int) bool
	dfs = func(l int) bool {
		if l == len(s) {
			return true
		}
		//if memo[l] == true {return true}
		if can, exist := memo[l]; exist {return can}
		for r := l; r < len(s) && r < l+maxL; r++ {
			if wd[s[l:r+1]] {
				//return dfs(r+1) // shouldn't directly return 
				if dfs(r+1) {
					memo[l] = true
					return true
				}
			}
		}
		// this line is very key, to skip the verified false case
		memo[l] = false
		return false
	}
	return dfs(0)
}
