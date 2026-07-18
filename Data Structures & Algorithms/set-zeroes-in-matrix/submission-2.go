func setZeroes(matrix [][]int) {
    r, c := len(matrix), len(matrix[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < r; k++ {
					if k != i && matrix[k][j] != 0 {
						matrix[k][j] = -1
					}
				}
				for k := 0; k < c; k++ {
					if k != j && matrix[i][k] != 0 {
						matrix[i][k] = -1
					}
				}
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
}
