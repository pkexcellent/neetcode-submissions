func lastStoneWeightII(stones []int) int {
	// this is basically a knapsap bag problem
	// we firstly get the half sum of the stones
	// and try to reach to the half as close as possible
	// so we pick one or not pick this one, to get the running sum
	
	n := len(stones)

	total := 0
	for i := 0; i < n; i++ {
		total += stones[i]
	}
	target := total/2

	memo := make(map[[2]int]int)
	var dfs func(idx int, pickedSum int) int
	dfs = func(idx int, pickedSum int) int {
		if pickedSum > target || idx == n {
			return abs(pickedSum - (total - pickedSum))
		}
		if v, exist := memo[[2]int{idx, pickedSum}]; exist {
			return v
		}
		// pick this one or not, return the min one
		rs :=  min(dfs(idx+1, pickedSum + stones[idx]),
					dfs(idx+1, pickedSum))
		memo[[2]int{idx, pickedSum}] = rs
		return rs
	}
	return dfs(0, 0)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
