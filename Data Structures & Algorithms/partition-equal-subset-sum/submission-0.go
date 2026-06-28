func canPartition(nums []int) bool {
    // total/2 is the target
	// then like bag problem
	sum := 0 
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum/2

	// row is number i, col is sum
	dp := make([][]bool, len(nums))
	for i, _ := range dp {
		dp[i] = make([]bool, target+1)
	}
	//if nums[0] > target {return false}
	for i := 0; i < len(nums); i++ {
		dp[i][0] = true // using first i numbers to sum at 0, is always true
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j < target+1; j++ {
			// pick or not pick
			// must nums[i] < j, then
			if j-nums[i] >=0 {
				// pick or not pick nums[i]
				dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}
