func getConcatenation(nums []int) []int {
    /*
    var n = len(nums)
    var ans = make([]int, 2*n)
    for i := 0; i < n; i ++ {
        ans[i] = nums[i]
        ans[i+n] = nums[i]
    }
    return ans
    */
    return append(nums, nums...)
}
