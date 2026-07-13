func canJump(nums []int) bool {
    // use a dp[i] indicate in ith, how many potential steps remain
    // so dp[i] = max(dp[i-1]-1, nums[i])
    n := len(nums)
    if n == 0 {
        return true
    }
    dp := make([]int, n)
    dp[0] = nums[0]
    for i := 1; i < n; i++ {
        if dp[i-1] <= 0 {return false}
        dp[i] = max(dp[i-1]-1, nums[i])
        if dp[i] == 0 && i != n-1 {
            return false
        }
    }
    return true
}
