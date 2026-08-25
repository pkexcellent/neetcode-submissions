func makesquare(matchsticks []int) bool {
	// calculate square's side
	total := 0
	for _, matchstick := range matchsticks {
		total += matchstick
	}
	if total%4 != 0 {
		return false
	}
	side := total/4
	// group them into 4*side
	sort.Slice(matchsticks, func(i, j int) bool {return matchsticks[i] > matchsticks[j]})
	
	sides := make([]int, 4)
	var dfs func(idx int) bool 
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			if sides[0] == side && sides[1] == side && 
				sides[2] == side && sides[3] == side {
					return true
				}
			return false
		}
		for i := 0; i < 4; i++ {
			if sides[i] + matchsticks[idx] > side {
				continue
			}
			if i > 0 && sides[i] == sides[i-1] {
				continue
			}
			sides[i] += matchsticks[idx]
			if dfs(idx+1) == true {
				return true
			}
			sides[i] -= matchsticks[idx]
		}
		return false
	}
	
	return dfs(0)
	
}
