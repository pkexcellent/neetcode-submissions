func lengthOfLIS(nums []int) int {
    // use bottom-up
	dp := make([]int, len(nums)+1)
	maxL := 1
	for i := len(nums)-1; i >=0; i-- {
		// just fill nums[i]
		dp[i] = 1
		for j := i+1; j < len(nums); j++ {
			// try to fill nums[i]
			if nums[i] < nums[j] {
				//fmt.Println(i, j, dp[j] + 1, dp[i])
				dp[i] = max(dp[j] + 1, dp[i])
			}
		}
		maxL = max(maxL, dp[i])
	}
	//fmt.Println(dp)
	return maxL
}
