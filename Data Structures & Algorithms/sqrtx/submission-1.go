func mySqrt(x int) int {
	// binary try again?
	// starting try between 1 and x/2
	l, r := 1, x
	for l <= r {
		m := l + (r-l)/2
		if m*m > x {
			r = m - 1
		} else if m*m < x {
			l = m + 1
		} else {
			return m
		}
	}
	return r
}
