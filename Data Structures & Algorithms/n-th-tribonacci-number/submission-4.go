func tribonacci(n int) int {
	initSet := make([]int, 3)
	initSet[0] = 0
	initSet[1] = 1
	initSet[2] = 1
	if n < 3 {
		return initSet[n]
	}
	for i := 3; i <= n; i++ {
		tmp := initSet[0] + initSet[1] + initSet[2]
		initSet[0] = initSet[1]
		initSet[1] = initSet[2]
		initSet[2] = tmp
	}
	return initSet[2]
}
