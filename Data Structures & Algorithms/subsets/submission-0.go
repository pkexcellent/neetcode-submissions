func subsets(nums []int) [][]int {
    var rs = [][]int{}
    var set = []int{}

    var dfs func(int)
    dfs = func(i int) {
        if i >= len(nums) {
            tmp := make([]int, len(set))
            copy(tmp, set)
            rs = append(rs, tmp)
            return
        }
        set = append(set, nums[i])
        dfs(i+1)
        set = set[0:len(set)-1]
        dfs(i+1)
    }
    dfs(0)
    return rs

}
