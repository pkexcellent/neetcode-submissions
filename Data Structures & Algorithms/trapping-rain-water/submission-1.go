func trap(height []int) int {
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	maxSoFar := 0
	for i := 0; i < n; i++ {
		if height[i] > maxSoFar {
			maxSoFar = height[i]
		}
		leftMax[i] = maxSoFar
	}
	maxSoFar = 0
	for i := n-1; i >= 0; i-- {
		if height[i] > maxSoFar {
			maxSoFar = height[i]
		}
		rightMax[i] = maxSoFar
	}
	
	total := 0
	for i := 0; i < n; i++ {
		barry := min(leftMax[i], rightMax[i])
		if height[i] < barry {
			total += barry - height[i]
		}
	}
	return total
}
