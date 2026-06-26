func wordBreak(s string, wordDict []string) bool {
    // bottom-up
	wd := make(map[string]bool)
	maxL := 0
	for _, w := range wordDict {
		wd[w] = true
		if len(w) > maxL {maxL = len(w)}
	}
	// states
	memo := make(map[int]bool, len(s)+1)
	memo[len(s)] = true
	for r := len(s); r >= 0; r-- {
		if memo[r] == false {
			continue
		}
		for l := r-1; l >= 0 && l >= r-maxL; l-- {
			if wd[s[l:r]] {
				memo[l] = true
			}
		}
	}
	return memo[0]
}
