func minInterval(intervals [][]int, queries []int) []int {
    // use sweep line to check which intervals can be used for each query
	// and then pick the shortest interval
	ninterval, nqueries := len(intervals), len(queries)
	events := make([][]int, ninterval*2 + nqueries)
	for i, interval := range intervals {
		start, end := interval[0], interval[1]
		length := end-start+1
		events[i*2] = []int{start, 1, length, i}
		events[i*2+1] = []int{end, -1, length, i}
	}
	for i, query := range queries {
		events[ninterval*2 + i] = []int{query, 0, i}
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] > events[j][1]
		}
		return events[i][0] < events[j][0]
	})
	//fmt.Println(events)
	rs := make([]int, nqueries)
	availableInt := make(map[int]int)

	for i := 0; i < len(events); i++ {
		if events[i][1] == 0 {
			queryNum := events[i][2]
			rs[queryNum] = selectLongest(availableInt)
			//fmt.Println("result: ", availableInt, rs)
		} else if events[i][1] == 1 {
			// start interval
			intervalIdx := events[i][3]
			availableInt[intervalIdx] = events[i][2]
		} else if events[i][1] == -1 {
			// end of interval
			intervalIdx := events[i][3]
			delete(availableInt, intervalIdx)
		}
		//fmt.Println("one round: ", availableInt)
	}
	//fmt.Println(rs)
	return rs
}
func selectLongest(mp map[int]int) int {
	if len(mp) == 0 {return -1}
	minV := math.MaxInt32
	for _, v := range mp {
		if v < minV {
			minV = v
		}
	}
	return minV
}
