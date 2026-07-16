func eraseOverlapIntervals(intervals [][]int) int {
    // greedy
	cnt := 0
	sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})
	prevE := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		curS := intervals[i][0]
		curE := intervals[i][1]
		if curS < prevE {
			// has overlap
			if curE < prevE {
				prevE = curE
			} // else do nothing, keep the previous
			cnt++
		} else { // no overlap, select this one
			prevE = curE
		}
	}
	return cnt
}
