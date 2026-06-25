func solveNQueens(n int) [][]string {
	// use dfs, there are N Qs, so one row for each
	rs := [][]string{}
	// one set of the final result, but we want to use []string to place Q
	// instead of using the final string like "...Q..."
	one := make([][]string, n)
	for i := range one {
		one[i] = make([]string, n)
		for c := 0; c < n; c++ {
			one[i][c] = "."
		}
	}

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			oneResult := []string{}
			for _, line := range one {
				tmp := ""
				for _, c := range line {
					tmp += c
				}
				oneResult = append(oneResult, tmp)
			}
			rs = append(rs, oneResult)
		}
		for column := 0; column < n; column++ {
			if isSafe(row, column, one) {
				one[row][column] = "Q"
				dfs(row+1)
				one[row][column] = "."
			}
		}
	}
	dfs(0)
	return rs
}

func isSafe(row, column int, board [][]string) bool {
	// check the current row
	//for i := 0; i < ;
	// check the current column
	for i := row-1; i >= 0; i-- {
		if board[i][column] == "Q" {return false}
	}
	// check corner left
	for r, c := row-1, column-1; r>=0 && c>=0; r, c = r-1, c-1 {
		if board[r][c] == "Q" {return false}
	} 

	for r, c := row-1, column+1; r>=0 && c<len(board); r, c = r-1, c+1 {
		if board[r][c] == "Q" {return false}
	} 
	return true
}

