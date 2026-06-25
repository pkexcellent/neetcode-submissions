func numIslands(grid [][]byte) int {
    var cnt = 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	var dfsMark func(r, c int) 
	dfsMark = func(r, c int) {
		visited[r][c] = true
		if r+1 < len(grid) && visited[r+1][c] == false && grid[r+1][c] == '1' {
			dfsMark(r+1, c)
		}
		if r-1 >= 0 && visited[r-1][c] == false && grid[r-1][c] == '1' {
			dfsMark(r-1, c)
		}
		if c-1 >= 0 && visited[r][c-1] == false && grid[r][c-1] == '1' {
			dfsMark(r, c-1)
		}
		if c+1 < len(grid[0]) && visited[r][c+1] == false && grid[r][c+1] == '1' {
			dfsMark(r, c+1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && visited[i][j] == false {
				cnt++
				dfsMark(i, j)
			}
		}
	}
	return cnt
}
