func eraseOverlapIntervals(intervals [][]int) int {
    // use DFS
	// dfs is to try to pick as more intervals as possible
	// when picking, we can choose whether pick it, or not
	// if pick it, we need to make sure there is no interval available now

	sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})
	// memo to store status
	memo := make(map[[2]int]int)
	var dfs func(cur, prev int) int
	dfs = func(cur, prev int) int {
		if cur == len(intervals) {
			return 0 // no one can pick
		}
		if v, exist := memo[[2]int{cur, prev}]; exist {
			return v
		}
		// not pick this one
		rs := dfs(cur+1, prev)
		// pick this one, try to
		if prev == -1 || (intervals[prev][1] <= intervals[cur][0]) {
			rs = max(rs, 1+dfs(cur+1, cur))
		}
		memo[[2]int{cur, prev}] = rs
		return rs
	}
	return len(intervals) - dfs(0, -1)

}
