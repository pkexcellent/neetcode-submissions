func partition(s string) [][]string {
	patterns := [][]string{}
	set := []string{}

	var dfs func(int, int)
	dfs = func(i, j int) {
		if j == len(s) {
			if i == j {
				// only when i==j, it means the remaining part is palindorome
				// only when this case, becasue of dfs(j+1, j+1), i==j, otherwise
				// the last piece is not palindorome, so we shouldn't put it into the final result
				// we should only return
				tmp := append([]string{}, set...)
				patterns = append(patterns, tmp)
			}
			return
		}
		if isPalindorome(s, i, j) {
			// continus to dfs, startring from j+1
			set = append(set, s[i:j+1])
			dfs(j+1, j+1)
			// backtracking the last one
			set = set[:len(set)-1]
		}
		// without puting the posibility
		dfs(i, j+1)
	}
	dfs(0, 0)
	return patterns
}

func isPalindorome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
