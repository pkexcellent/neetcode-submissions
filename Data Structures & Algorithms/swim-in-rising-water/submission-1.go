func swimInWater(grid [][]int) int {
    // binary search + DFS
	// get the max and min heights in the grid
	maxH, minH := grid[0][0], grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			maxH = max(maxH, grid[i][j])
			minH = min(minH, grid[i][j])
		}
	}

	// dfs to find the path with a given time t <- means the max height is t
	visited := make([][]bool, len(grid))
	for i, _ := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	var dfs func(maxHeight int, r, c int) bool
	dfs = func(maxHeight int, r, c int) bool {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] > maxHeight || visited[r][c] {
			return false
		}
		if r == len(grid) - 1 && c == len(grid) - 1 {
			return true
		}
		visited[r][c] = true
		return dfs(maxHeight, r-1, c) || dfs(maxHeight, r+1, c)||
		 dfs(maxHeight, r, c-1) || dfs(maxHeight, r, c+1)
	}

	// binary - to reduce looping every possible t
	l, r := minH, maxH
	for l < r {
		timeToHeight := (l+r)/2
		if dfs(timeToHeight, 0, 0) {
			r = timeToHeight
		} else {
			l = timeToHeight + 1
		}
		// mark visited as false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				visited[i][j] = false
			}
		}
	}
	return r
}
