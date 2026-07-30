type NumMatrix struct {
	content [][]int
	prefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	prefix := make([][]int, m+1)
	for i, _ := range prefix {
		prefix[i] = make([]int, n+1)
		prefix[i][0] = 0
	}
	for i := 0; i < n+1; i++ {
		prefix[0][i] = 0
	}
	for i := 1; i < m+1; i++ {
		rowSum := 0
		for j := 1; j < n+1; j++ {
			rowSum += matrix[i-1][j-1]
			prefix[i][j] = prefix[i-1][j] + rowSum
		}
	}
	return NumMatrix{
		content: matrix,
		prefixSum: prefix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.prefixSum[row2+1][col2+1] - 
	this.prefixSum[row2+1][col1] - 
	this.prefixSum[row1][col2+1] + 
	this.prefixSum[row1][col1]
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
