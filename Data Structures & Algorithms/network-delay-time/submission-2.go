func networkDelayTime(times [][]int, n int, k int) int {
	// build adjacent table
	adj := make(map[int][][2]int)
	for _, time := range times {
		src, dst, t := time[0], time[1], time[2]
		adj[src] = append(adj[src], [2]int{dst, t})
	}

	//fmt.Println(adj)
	visited := make(map[int]bool, n)
	q := [][2]int{}
	q = append(q, adj[k]...)
	visited[k] = true
	t := 0
	for len(q) > 0 {
		layerSize := len(q)
		t++
		//fmt.Println(q)
		for i := 0; i< layerSize; i++ {
			dst, remainingT := q[0][0], q[0][1]
			q = q[1:]
			if visited[dst] {
				continue
			}
			// only when dst t is 1, this is reachable
			if remainingT <= 1 {
				visited[dst] = true
				q = append(q, adj[dst]...)
			} else { // else t-1, and put back to the queue again
				q = append(q, [2]int{dst, remainingT-1})
			}
		}
		if len(visited) == n {
			return t
		}
	}
	return -1
}
