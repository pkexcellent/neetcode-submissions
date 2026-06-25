func islandsAndTreasure(grid [][]int) {
    // BFS is a very nature mathod for calculating shortest distance

	q := [][2]int{}
	for i := 0; i < len(grid); i++ {
		for j:=0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		x, y := cur[0], cur[1]
		if x-1 >= 0 && grid[x-1][y] == 2147483647 {
			grid[x-1][y] = grid[x][y] + 1
			q = append(q, [2]int{x-1, y})
		}
		if y-1 >= 0 && grid[x][y-1] == 2147483647 {
			grid[x][y-1] = grid[x][y] + 1
			q = append(q, [2]int{x, y-1})
		}
		if x+1 < len(grid) && grid[x+1][y] == 2147483647 {
			grid[x+1][y] = grid[x][y] + 1
			q = append(q, [2]int{x+1, y})
		}
		if y+1 < len(grid[0]) && grid[x][y+1] == 2147483647 {
			grid[x][y+1] = grid[x][y] + 1
			q = append(q, [2]int{x, y+1})
		}
	}
}
