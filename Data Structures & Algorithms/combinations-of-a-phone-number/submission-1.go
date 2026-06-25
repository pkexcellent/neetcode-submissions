func letterCombinations(digits string) []string {
	// use dfs to solve
	if len(digits) == 0 {
		return []string{}
	}
	digitMap := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
	}
	rs := []string{}

	var dfs func(int, string)
	dfs = func(i int, curString string) {
		if i == len(digits) {
			rs = append(rs, curString)
			return
		}
		for _, c := range digitMap[digits[i]] {
			dfs(i+1, curString+string(c))
		}
	}
	dfs(0, "")
	return rs
}
