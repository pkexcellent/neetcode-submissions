func maxCoins(nums []int) int {
    // dp[l][r] means the max coins between l and r
    // so by bursting ballon in i, if i is the last ballon to burst, then
    // the coin is dp[l][i-1] + (nums[l-1]*nums[i]*nums[r+1]) + dp[i+1][r]
    // so the final max dp[l][r] is the max one among bursting every ballon i between l and r

    // put two 1 in head and tail
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)

    n := len(nums)
    dp := make([][]int, n)
    for i, _ := range dp {
        dp[i] = make([]int, n)
    }

    // not involving the first and last 1 we just inserted
    for l := n-2; l >= 1; l-- {
        for r := l; r <= n-2; r++ {
            maxCoins := 0
            for i := l; i <= r; i++ {
                coins := dp[l][i-1] + dp[i+1][r] + nums[l-1]*nums[i]*nums[r+1]
                if coins > maxCoins {
                    maxCoins = coins
                }
            }
            dp[l][r] = maxCoins
        }
    }
    return dp[1][n-2]
}
