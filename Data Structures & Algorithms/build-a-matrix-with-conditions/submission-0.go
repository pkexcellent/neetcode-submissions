func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	// build a sorted list for row and col
	// using topology sort
	rowDist, colDist := make(map[int]int), make(map[int]int)
	rowOrder, colOrder := getOrderedList(rowConditions, k), getOrderedList(colConditions, k)
	fmt.Println(rowOrder, colOrder)
	if len(rowOrder) == 0 || len(colOrder) == 0 {
		return [][]int{}
	}
	for i, num := range rowOrder {
		rowDist[num] = i
	}
	for i, num := range colOrder {
		colDist[num] = i
	}
	rs := make([][]int, k)
	for i, _ := range rs {
		rs[i] = make([]int, k)
	}
	for i := 1; i <= k; i++ {
		row, col := rowDist[i], colDist[i]
		rs[row][col] = i
	}
	return rs
}

func getOrderedList(edges [][]int, n int) []int {
	adj := make(map[int][]int)
	indegrees := make(map[int]int)
	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		adj[src] = append(adj[src], dst)
		indegrees[dst]++
	}
	q := []int{}
	rs := []int{}
	for i := 1; i <= n; i++ {
		if _, exist := indegrees[i]; !exist {
			q = append(q, i)
			rs = append(rs, i)
		}
	}
	fmt.Println(adj, indegrees)

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			first := q[0]
			q = q[1:]
			for _, dst := range adj[first] {
				indegrees[dst]--
				if indegrees[dst] == 0 {
					q = append(q, dst)
					rs = append(rs, dst)
				}
			}
		}
	}
	return rs
}
