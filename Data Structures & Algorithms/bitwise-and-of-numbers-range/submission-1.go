func rangeBitwiseAnd(left int, right int) int {
	// shit, get the common prefix
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}
