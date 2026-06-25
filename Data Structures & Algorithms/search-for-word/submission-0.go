func exist(board [][]byte, word string) bool {
	var x = len(board)
	var y = len(board[0])
	var visited = make([][]bool, x)
	for i := range visited {
		visited[i] = make([]bool, y) 
	}

	var dfs func(int, int, int) bool
	dfs = func(i, j, index int) bool {
		if index == len(word) {
			return true
		} 
		if i < 0 || i>=x || j<0 || j>=y || board[i][j] != word[index] || visited[i][j] == true {
			return false
		}
		visited[i][j] = true
		rs := dfs(i+1, j, index+1) || 
			dfs(i-1, j, index+1) ||
			dfs(i, j+1, index+1) ||
			dfs(i, j-1, index+1)
		visited[i][j] = false
		return rs
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if dfs(i, j, 0) == true {
				return true
			}
		}
	}
	return false

}
