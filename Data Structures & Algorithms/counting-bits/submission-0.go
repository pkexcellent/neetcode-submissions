func countBits(n int) []int {
	// the continues 1 in low bits is x, then + 1 mean reduce x, but +1
	rs := make([]int, n+1)
	rs[0] = 0
	for i := 1; i < n+1; i++ {
		//low 1s
		lowOnes := 0
		cur := i-1
		for cur & 1 > 0 {
			lowOnes++
			cur = cur >> 1
		}

		rs[i] = rs[i-1] + 1 - lowOnes // change the last 1s to 0, and add a 1
	}
	return rs
}
