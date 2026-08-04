func largestRectangleArea(heights []int) int {
	// monotonic stack
	stack := []int{}
	maxs := 0
	// need to push a zero in rear, the 0 is to trigger pop out the stack
	heights = append(heights, 0)
	for i := 0; i< len(heights); i++ {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				// 出去的一定是单调的顺序，所以w的右边界就是i（i的左侧）
				// w的左边界就是左边挨着它比他矮的第一个右侧
				// 所以w 是 (i-1)-l 或者i- (l+1)，反正要用左侧就都用左侧，用右侧就都用右侧
				top := heights[stack[len(stack)-1]]
				stack = stack[:len(stack) - 1]
				w := i // 记住这里的i是当前这个height的右侧一个，所以需要是i-1
				if len(stack) > 0 {
					w = i - stack[len(stack)-1] - 1
				}				
				h := top
				maxs = max(maxs, h*w)
			}
			stack = append(stack, i)
		}
	}
	return maxs
}
