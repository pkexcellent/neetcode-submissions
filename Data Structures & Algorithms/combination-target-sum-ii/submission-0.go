func combinationSum2(candidates []int, target int) [][]int {
    var rs = [][]int{}

    sort.Ints(candidates)

    var dfs func(int, []int, int)
    dfs = func(i int, set []int, total int) {
        if total == target {
            tmp := make([]int, len(set))
            copy(tmp, set)
            rs = append(rs, tmp)
            return
        }
        if i >= len(candidates) || total > target {
            return
        }
        // put candidates[i] into set
        set = append(set, candidates[i])
        dfs(i+1, set, total + candidates[i])
        // backtrack candidates[i]
        set = set[0:len(set)-1]
        // not put it into set
        // if we choose this way, we can safely skip the other "not put" one, becasue
        // they are the same
        for i+1 < len(candidates) && candidates[i] == candidates[i+1] {i++}
        dfs(i+1, set, total)
    }
    dfs(0, []int{}, 0)
    return rs
}
