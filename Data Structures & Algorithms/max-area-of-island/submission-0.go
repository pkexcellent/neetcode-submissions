func maxAreaOfIsland(grid [][]int) int {
    maxSize := 0

	visited := make([][]bool, len(grid))
	for i, _ := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	var dfsIsland func(int, int, *int)
	dfsIsland = func(r, c int, size *int) {
		visited[r][c] = true
		*size = *size + 1
		if r + 1 < len(grid) && !visited[r+1][c] && grid[r+1][c] == 1 {
			dfsIsland(r+1, c, size)
		}
		if r -1 >= 0 && !visited[r-1][c] && grid[r-1][c] == 1 {
			dfsIsland(r-1, c, size)
		}
		if c-1 >= 0 && !visited[r][c-1] && grid[r][c-1] == 1 {
			dfsIsland(r, c-1, size)
		}
		if c+1 < len(grid[0]) && !visited[r][c+1] && grid[r][c+1] == 1 {
			dfsIsland(r, c+1, size)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j< len(grid[i]); j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				thisSize := 0
				dfsIsland(i, j, &thisSize)
				if thisSize > maxSize {maxSize = thisSize}
			}
		}
	}
	return maxSize
}
