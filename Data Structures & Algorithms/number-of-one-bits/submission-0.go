func hammingWeight(n int) int {
	cnt := 0
	for n > 0 {
		if n & 1 > 0 {
			cnt += 1
		}
		n = n >> 1
	}
	return cnt
}
