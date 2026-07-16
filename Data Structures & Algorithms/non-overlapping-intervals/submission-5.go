func eraseOverlapIntervals(intervals [][]int) int {
    // bottom-up method
	// dp[cur] = max(dp[cur], 1 + all the dp[prev]s if prev can be placed before cur)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	dp := make([]int, len(intervals))
	maxInters := 0
	for i := 0; i < len(intervals); i++ {
		// place intervals[i]
		dp[i] = 1
		// try to place previous ones
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		maxInters = max(maxInters, dp[i])
	}
	return len(intervals) - maxInters
}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
