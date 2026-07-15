func merge(intervals [][]int) [][]int {
    // sort by start
    sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})

    rs := [][]int{}
    a := intervals[0]
    for i := 1; i < len(intervals); i++ {
        mg := tryMergeTwo(a, intervals[i])
        if len(mg) == 1 {
            a = mg[0]
        } else {
            rs = append(rs, mg[0])
            a = mg[1]
        }
    }
    rs = append(rs, a)
    return rs
}

func tryMergeTwo(a, b []int) [][]int {
    // assume a and b are already sorted by start
    // so a[0] must <= b[0]
    // so overlap == a[1] >= b[0] 
    if a[1] >= b[0] {
        return append([][]int{}, []int{min(a[0], b[0]), max(a[1], b[1])})
    } else {
        return append(append([][]int{}, a), b)
    }
}
