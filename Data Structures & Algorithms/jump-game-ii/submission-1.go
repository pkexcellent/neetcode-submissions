func jump(nums []int) int {
    // use dp[i] to store how many the min steps can reach i
    // repeatly update dp[i]
    n := len(nums)
    dp := make([]int, len(nums))
    for i := 0; i < n; i++ {
        dp[i] = math.MaxInt32
    }
    dp[0] = 0
    for i := 0; i < n; i++ {
        longest := i+nums[i]
        minStepAti := dp[i]
        for j := i; j < min(longest+1, n); j++ {
            dp[j] = min(dp[j], minStepAti+1) // i also means the ith jump
        }
        //fmt.Println(dp)
    }
    
    return dp[n-1]
}
