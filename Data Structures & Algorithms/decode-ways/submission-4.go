func numDecodings(s string) int {
    // dfs
	mem := make([]int, len(s))
	for i := 0; i< len(mem); i++ {
		mem[i] = math.MaxInt
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= len(s) {
			return 1
		} 
		if s[i] == '0' {
			return 0
		}
		if mem[i] != math.MaxInt {
			return mem[i]
		}
		countSoFar := 0
		countSoFar += dfs(i+1)
		if i+1 < len(s) {
			if s[i] == '1' || (s[i] == '2' && s[i+1] < '7') {
				countSoFar += dfs(i+2)
			}
		}
		mem[i] = countSoFar
		return countSoFar
	}
	return dfs(0)
}
