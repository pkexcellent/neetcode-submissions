func insert(intervals [][]int, newInterval []int) [][]int {
    // greedy
    // 3 cases, the order of the conditions is very important
    // 1. (not correct) newInterval-start > interval-end: append interval
    // should be 1. newInterval-end < interval-start: append new
    //  else 2. newInterval-start > interval-end: append interval
    // otherwise, merge
    rs := [][]int{}

    for i := 0; i < len(intervals); i++ {
        if newInterval[0] > intervals[i][1] {
            rs = append(rs, intervals[i])
        } else if newInterval[1] < intervals[i][0] {
            rs = append(rs, newInterval)
            return append(rs, intervals[i:len(intervals)]...)
        } else {
            newInterval[0] = min(newInterval[0], intervals[i][0])
            newInterval[1] = max(newInterval[1], intervals[i][1])
        }
    }
    rs = append(rs, newInterval)
    return rs
}
