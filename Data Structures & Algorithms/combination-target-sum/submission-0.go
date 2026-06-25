func combinationSum(nums []int, target int) [][]int {
    // it seems there isn't an easy way to avoid using backtrack
    var rs = [][]int{}
    
    var dfs func(int, []int, int)
    dfs = func(i int, set []int, total int) {
        // quit condition
        if i >= len(nums) || total > target {
            return
        }
        if target == total {
            tmp := make([]int, len(set))
            copy(tmp, set)
            rs = append(rs, tmp)
            return
        }
        // put num[i] into set
        //set = append(set, nums[i])
        dfs(i, append(set, nums[i]), total + nums[i])
        // backtrack num[i]
        //set = set[0:len(set)-1]

        //not put num[i] into set
        dfs(i+1, set, total)
    }
    dfs(0, []int{}, 0)
    return rs
}
