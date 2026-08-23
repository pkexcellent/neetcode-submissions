func canReach(s string, minJump int, maxJump int) bool {
	// use a simply DFS + memo
	n := len(s)
	if s[n-1] != '0' {return false}

	// memo
	able := make(map[int]bool, n)
	var dfs func(cur int) bool
	dfs = func(cur int) bool {
		if cur == n-1 {
			return true
		}
		if cache, exist := able[cur]; exist {
			return cache
		}
		rs := false
		for i := cur + minJump; i <= cur + maxJump && i < n; i++ {
			if s[i] == '0' {
				rs = rs || dfs(i)
			}
		}
		able[cur] = rs
		return rs
	}	
	rrss := dfs(0)
	fmt.Println(able)
	return rrss
}
