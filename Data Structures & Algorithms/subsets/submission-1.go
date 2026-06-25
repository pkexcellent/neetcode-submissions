func subsets(nums []int) [][]int {
    var rs = [][]int{}
    rs = append(rs, []int{})
    for _, num := range nums {
        size := len(rs)
        for i := 0; i < size; i++ {
            var newSet = make([]int, len(rs[i]))
            copy(newSet, rs[i])
            newSet = append(newSet, num)
            rs = append(rs, newSet)
        }
    }
    return rs
}
