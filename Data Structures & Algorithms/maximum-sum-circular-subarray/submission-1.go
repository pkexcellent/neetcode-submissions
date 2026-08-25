func maxSubarraySumCircular(nums []int) int {
	// either it's the max in middle - no wrap
	// or it's the sum - min in middle - has wrap
	n := len(nums)
	// get max
	subsum := 0
	maxSubSum := 0
	for i := 0; i < n; i++ {
		subsum += nums[i]
		if subsum < 0 {
			subsum = 0
		} else {
			maxSubSum = max(maxSubSum, subsum)
		}
	}
	subsum = 0
	minSubSum := 0
	maxNum := nums[0]
	for i := 0; i < n; i++ {
		maxNum = max(maxNum, nums[i])
		subsum += nums[i]
		if subsum > 0 {
			subsum = 0
		} else {
			minSubSum = min(minSubSum, subsum)
		}
	}

	totalSum := 0
	for i := 0; i < n; i++ {
		totalSum += nums[i]
	}

	if maxNum <= 0 {
		return maxNum
	} else {
		return max(maxSubSum, totalSum - minSubSum)
	}
	
}
