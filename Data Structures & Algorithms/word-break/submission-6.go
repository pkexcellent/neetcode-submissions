func wordBreak(s string, wordDict []string) bool {
    memo := make(map[int]bool, len(s)+1)
	memo[len(s)] = true

	for i := len(s)-1; i >= 0; i-- {
		for _, w := range wordDict {
			if i+len(w) <= len(s) {
				if s[i:i+len(w)] == w && memo[i+len(w)] == true {
					memo[i] = true
					break
				}
			}
		}
	}
	return memo[0]
}
