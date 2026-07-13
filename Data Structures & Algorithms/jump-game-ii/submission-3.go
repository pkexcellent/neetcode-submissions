func jump(nums []int) int {
    // BFS, use l, r to represent the current window
    // which is like a "level" when BFS
    // we calculate to farest, so then window can be updated to 
    // l = last_r+1, r = farest
    l, r := 0, 0
    step := 0
    for r < len(nums)-1 { // has to use len-1, becasue the last pos we don't need to use
        farest := r
        for i := l; i <= r; i++ {
            farest = max(farest, i + nums[i])
        }
        // update window
        l = r+1
        r = farest
        step++
    }
    return step
}
