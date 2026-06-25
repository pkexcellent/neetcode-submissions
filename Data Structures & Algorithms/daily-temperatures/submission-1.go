func dailyTemperatures(temperatures []int) []int {
	// use monotonic stack
	var ans = make([]int, len(temperatures))
	var stack = []int{} // store the index of temperatures
	for idx, temp := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, idx)
		} else {
			topIdx := stack[len(stack)-1]
			if temperatures[topIdx] >= temp {
				stack = append(stack, idx)
			} else {
				for temperatures[topIdx] < temp {
					ans[topIdx] = idx - topIdx
					stack = stack[:len(stack)-1]
					if len(stack) == 0 {
						break
					} else {
						topIdx = stack[len(stack)-1]
					}
				}
				stack = append(stack, idx)	
			}
		}
	}
	for _, idx := range stack {
		ans[idx] = 0
	}
	return ans
}
