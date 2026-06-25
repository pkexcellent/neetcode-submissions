func climbStairs(n int) int {
	if n <=2 {return n}
    step1, step2 := 1, 2
	for i := 2; i < n; i++ {
		tmp := step1+step2
		step1 = step2
		step2 = tmp
	}
	return step2
}
