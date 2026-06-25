func solve(board [][]byte) {
    // use DFS or bfs to scan cells that can be passed with boarder
	// then flip, flip

	var dfs func(r, c int) 
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) ||
			board[r][c] != 'O' {
				return
			}
		if board[r][c] == 'O' {
			board[r][c] = 'A'
		}
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][len(board[0])-1] == 'O'{
			dfs(i, len(board[0])-1)
		}
	}
	for j := 0; j< len(board[0]); j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[len(board)-1][j] == 'O' {
			dfs(len(board) -1, j)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
	return
}
