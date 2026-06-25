func subsets(nums []int) [][]int {
    // using bit
    var rs = [][]int{}

    for i := 0; i < (1<<len(nums)); i++ {
        set := []int{}
        for bit := 0; bit < len(nums); bit++ {
            if i & (1 << bit) > 0 {
                set = append(set, nums[bit])
            }
        }
        rs = append(rs, set)
    }
    return rs
}
