func pacificAtlantic(heights [][]int) [][]int {
    // it will exceed the time limit if DFS every point, even adding a tracker
	// so we have to dfs just from the edges
	// if from both edges, a point can be reached, then this point is the result

	var pacificAccess = make(map[[2]int]bool)
	var atlanticAccess = make(map[[2]int]bool)

	var dfs func(r, c int, preHeight int, visited map[[2]int]bool)
	dfs = func(r, c, preHeight int, visited map[[2]int]bool) {
		if r < 0 || c < 0 || r >= len(heights) || c >= len(heights[0]) || 
			visited[[2]int{r, c}] || heights[r][c] < preHeight {
				// note, this is heights[r][c] < preHeight because it's the reversed flow
				return
			}
			visited[[2]int{r, c}] = true
			dfs(r+1, c, heights[r][c], visited)
			dfs(r-1, c, heights[r][c], visited)
			dfs(r, c+1, heights[r][c], visited)
			dfs(r, c-1, heights[r][c], visited)
	}
	for i := 0; i < len(heights); i++ {
		dfs(i, 0, heights[i][0], pacificAccess)
		dfs(i, len(heights[0])-1, heights[i][len(heights[0])-1], atlanticAccess)
	}
	for j := 0; j < len(heights[0]); j++ {
		dfs(0, j, heights[0][j], pacificAccess)
		dfs(len(heights)-1, j, heights[len(heights)-1][j], atlanticAccess)
	}
	rs := [][]int{}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			point := [2]int{i, j}
			if pacificAccess[point] && atlanticAccess[point] {
				rs = append(rs, []int{i, j})
			}
		}
	}
	return rs
}
