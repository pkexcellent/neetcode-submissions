func countBits(n int) []int {
	rs := make([]int, n+1)
	rs[0] = 0
	for i := 1; i < n+1; i++ {
		cur := i
		one := 0 
		for cur > 0 {
			one++
			cur = cur & (cur-1)
		}
		rs[i] = one
	}
	return rs
}
