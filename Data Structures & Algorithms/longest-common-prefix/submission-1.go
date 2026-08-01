func longestCommonPrefix(strs []string) string {
    prefix := ""
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	i := 0
	cur := strs[0][i]
	for {
		for _, s := range strs {
			if i >= len(s) {
				return prefix
			}
			if s[i] != cur {
				return prefix
			}
		}
		prefix = prefix + string(cur)
		i++
		if i >= len(strs[0]) {
			return prefix
		}
		cur = strs[0][i]
	}
	return prefix
}
