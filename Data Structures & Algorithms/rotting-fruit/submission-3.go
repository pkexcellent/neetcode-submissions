func orangesRotting(grid [][]int) int {
    // push a minute while BFS
	totalFresh := 0
	q := [][2]int{}
	for i := 0; i< len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			}
			if grid[i][j] == 1 {
				totalFresh++
			}
		}
	}
	t := -1
	q = append(q, [2]int{-1, 1})

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(q) > 0 {
		cur := q[0]
		if cur[0] == -1 {
			// this is the time point
			t += cur[1]
			q = q[1:]
			if len(q) > 0 {
				cur = q[0]
				q = append(q, [2]int{-1, 1})
			} else {
				break
			}
		}
		q = q[1:]

		//hasNew := false
		for _, dir := range directions {
			r, c := cur[0] + dir[0], cur[1] + dir[1]
			if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != 1 {
				continue
			}
			//hasNew = true
			grid[r][c] = 2
			totalFresh--
			q = append(q, [2]int{r, c})
		}
	}
	if totalFresh > 0 {
		return -1
	}
	return t
}
