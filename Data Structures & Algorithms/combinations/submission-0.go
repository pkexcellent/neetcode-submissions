func combine(n int, k int) [][]int {
	// dfs
	rs := [][]int{}
	var dfs func(cur int, remaining int, picked []int) 
	dfs = func(cur int, remaining int, picked []int) {
		if remaining == 0 {
			tmp := make([]int, len(picked))
			copy(tmp, picked)
			rs = append(rs, tmp)
			return
		}
		if n - cur + 1 < remaining {
			return
		}
		// pick cur
		dfs(cur+1, remaining - 1, append(picked, cur))
		// not pick cur
		dfs(cur+1, remaining, picked)
	}
	dfs(1, k, []int{})
	return rs
}
