func totalNQueens(n int) int {
	table := make([][]byte, n)
	for i, _ := range table {
		table[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			table[i][j] = '.'
		}
	}
	rs := 0
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			rs++
			return
		}
		for col := 0; col < n; col++ {
			table[row][col] = 'Q'
			if isValid(table, row, col) {
				dfs(row+1)
			}
			table[row][col] = '.'
		}
	}
	dfs(0)
	return rs
}

func isValid(table [][]byte, row, col int) bool {
	// check col, no need to check row, because this is the only one
	n := len(table)
	for i := 0; i < row; i++ {
		if table[i][col] == 'Q' {
			return false
		}
	}
	// check left up cornor
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if table[r][c] == 'Q' {
			return false
		}
	}
	// check right up conor
	for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
		if table[r][c] == 'Q' {
			return false
		}
	}
	return true
}
