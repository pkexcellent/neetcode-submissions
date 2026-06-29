func uniquePaths(m int, n int) int {
    // use math method
	// pick n-1 from m-1+n-1
	// C n-1+m-1 m-1
	// ((n-1+m-1)*....*(n-1+m-1-n+1))/ (n*(n-1)*...*1)
	// m*(m-1)*....*(m-n+1) / n!, (m - (m-n+1) + 1) = n. n-1+1 = n
	// C 5 2 = 5*4/2*1
	// C 5 3 = 5*4*3/3*2*1
	x := min(n, m)
	res := 1
	for i, j := (n-1)+(m-1), 1; i >= (n-1)+(m-1)-(x-1)+1; i, j = i-1, j+1 {
		res = res*i/j
	}

	return res
}
