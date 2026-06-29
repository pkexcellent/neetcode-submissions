func canPartition(nums []int) bool {
    total := 0
	for _, num := range nums {
		total += num
	}
	if total%2 != 0 {
		return false
	}
	target := total/2

	// use a dp to mark which sum can be reached
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j-num >= 0; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[target]
}
