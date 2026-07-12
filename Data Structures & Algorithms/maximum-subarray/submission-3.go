func maxSubArray(nums []int) int {
    // starting from a positive number, keep adding until the sum is negtive, 
    // then restart from a positive number
    // maxSum
    // curSum: starting from a positive number

    if len(nums) == 0 {
        return 0
    }
    curSum, maxSum := nums[0], nums[0]
    curMax := curSum
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i] > 0 {
            if curSum < 0 {
                curSum = 0 // if the current sum is negative, restart from this number
            }
            curSum = curSum + nums[i]
            if curSum > curMax {
                curMax = curSum
            }
        } else if nums[i] + curSum > 0 {
            // although nums[i] is negative, but sum is positive, we still sum, but not update curMax
            curSum = curSum + nums[i]
        } else { // nums[i] is negative, sum is negative. break the continue sum, use nums[i] as start point
            curSum = nums[i]
        }
        maxSum = max(curSum, maxSum)
    }
    return maxSum
}
