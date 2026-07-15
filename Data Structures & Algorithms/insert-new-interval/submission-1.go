func insert(intervals [][]int, newInterval []int) [][]int {
   rs := [][]int{}
   curIdx := 0
   for curIdx < len(intervals) {
        if intervals[curIdx][1] < newInterval[0] {
            rs = append(rs, intervals[curIdx])
            curIdx++
        } else {
            break
        }
   }
   mergedS, mergedE := newInterval[0], newInterval[1]
   for curIdx < len(intervals) {
        if intervals[curIdx][0] > newInterval[1] {
            break
        } else {
            mergedS = min(intervals[curIdx][0], mergedS)
            mergedE = max(intervals[curIdx][1], mergedE)
            curIdx++
        } 
   }
   rs = append(rs, []int{mergedS, mergedE})
   rs = append(rs, intervals[curIdx:len(intervals)]...)
   return rs
}
