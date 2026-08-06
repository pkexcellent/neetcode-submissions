// Split Array Largest Sum
/*
dp[i][j] = minimum possible maximum subarray sum when splitting nums[0..i-1] into j partitions
Base case: dp[i][1] = prefix[i] (sum of first i elements)
Transition: dp[i][j] = min over all m from j-1 to i-1 of max(dp[m][j-1], prefix[i] - prefix[m])
The answer is dp[n][k]
The rewrite uses bottom-up DP — O(n²k) time, O(nk) space:

prefix[i] = sum of nums[0..i-1]
dp[i][j] = minimum possible max-subarray-sum when splitting nums[0..i-1] into j parts
Transition: for each split point m, the last part is nums[m..i-1] with sum prefix[i]-prefix[m], and the first j-1 parts are already optimal in dp[m][j-1]
Answer: dp[n][k]
*/
func splitArray(nums []int, k int) int {
	n := len(nums)

	// prefix[i] = sum of nums[0..i-1]
	prefix := make([]int, n+1)
	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}

	// dp[i][j] = min possible max-subarray-sum splitting nums[0..i-1] into j parts
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = 1000000000
		}
	}
	dp[0][0] = 0

	for j := 1; j <= k; j++ {
		for i := j; i <= n; i++ {
			// try all split points m: first j-1 parts cover nums[0..m-1], last part is nums[m..i-1]
			for m := j - 1; m < i; m++ {
				dp[i][j] = min(dp[i][j], max(dp[m][j-1], prefix[i]-prefix[m]))
			}
		}
	}

	return dp[n][k]
}
