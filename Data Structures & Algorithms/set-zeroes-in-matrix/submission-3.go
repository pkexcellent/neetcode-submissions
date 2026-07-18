func setZeroes(matrix [][]int) {
    // solution1, set numbers to -1, or others all has posiblity to overwrite other numbers
	// so using the first column and first row can solve that
	isRow, isCol := false, false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if j == 0 {
					isCol = true
				}
				if i == 0 {
					isRow = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j ++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if isRow {
		for j := 1; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if isCol {
		for i := 1; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
