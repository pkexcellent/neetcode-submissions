func generateParenthesis(n int) []string {
	var rs = []string{}

	var backtrack func(int, int, int, string)

	backtrack = func(i int, left int, totalLeft int, subset string) {
		if i == 2*n {
			rs = append(rs, subset)
			return
		}
		// not put right, put left, if there is still available left
		if totalLeft < n {
			backtrack(i+1, left+1, totalLeft + 1, subset + "(")
		}
		// if left > 0, put right
		if left > 0 {
			backtrack(i+1, left-1, totalLeft, subset + ")")
		}
	}
	backtrack(0, 0, 0, "")
	return rs
}
