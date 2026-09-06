func tribonacci(n int) int {
	l := max(3, n+1)
	initSet := make([]int, l)
	initSet[0] = 0
	initSet[1] = 1
	initSet[2] = 1
	if n < 3 {
		return initSet[n]
	}
	for i := 3; i <= n; i++ {
		initSet[i] = initSet[i-1] + initSet[i-2] + initSet[i-3]
	}
	return initSet[n]
}
