func spiralOrder(matrix [][]int) []int {
	rs := []int{}
	down, right := len(matrix)-1, len(matrix[0])-1
	left, up := 0, 0
	for left <= right && up <= down {
			// fmt.Println("l=",left, "r=",right, "u=",up, "d=",down)
			rs = append(rs, readCircle(matrix, left, right, up, down)...)
			up++
			right--
			left++
			down--
	}
	return rs
}

func readCircle(matrix [][]int, l, r, u, d int) []int {
	rs := []int{}
	for i := l; i <= r; i++ {
		rs = append(rs, matrix[u][i])
	}
	for i := u+1; i <= d; i++ {
		rs = append(rs, matrix[i][r])
	}
	if d > u {
		for i := r-1; i >= l; i-- {
			rs = append(rs, matrix[d][i])
		}
	}
	if r > l {
		for i := d-1; i >= u+1; i-- {
			rs = append(rs, matrix[i][l])
		}
	}
	return rs
}
