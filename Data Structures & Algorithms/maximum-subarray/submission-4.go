func maxSubArray(nums []int) int {
    // same as solution1, but simplified 
    if len(nums) == 0 {
        return 0
    }
    localMax := nums[0]
    globalMax := localMax

    for i := 1; i < len(nums); i++ {
        if localMax + nums[i] > nums[i] {
            localMax = localMax + nums[i]
        } else {
            localMax = nums[i]
        }
        globalMax = max(globalMax, localMax)
    }
    return globalMax
}
