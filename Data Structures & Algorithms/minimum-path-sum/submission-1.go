func minPathSum(grid [][]int) int {
	// can use dfs or dp
	// first dfs
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = math.MaxInt32
		}
	}
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row == m-1 && col == n-1 {
			return grid[row][col]
		}
		if row >= m || col >= n {
			return math.MaxInt32
		}
		if memo[row][col] != math.MaxInt32 {
			return memo[row][col]
		}
		pathSum := grid[row][col] + min(dfs(row+1, col), dfs(row, col+1))
		memo[row][col] = pathSum
		return pathSum
	}
	return dfs(0, 0)
}
