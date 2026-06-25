func threeSum(nums []int) [][]int {
    var m = make(map[int][]int, len(nums))
    for idx, num := range nums {
        m[num] = append(m[num], idx)
    }
    var ans = [][]int{}
    var mapDedup = make(map[[3]int]bool)
    for idx, num := range nums {
        remainTwoSum := 0 - num
        for i := idx+1; i<len(nums); i++ {
            lastNum := remainTwoSum - nums[i]
            if lastIdx, exist := m[lastNum]; exist {
                var triplet = []int{nums[idx], nums[i], nums[lastIdx[0]]}
                sort.Ints(triplet)
                var lastIdxOfLast = lastIdx[len(lastIdx)-1]
                if lastIdxOfLast > i {
                    if _, exist := mapDedup[[3]int(triplet)]; !exist {
                        ans = append(ans, triplet)
                        mapDedup[[3]int(triplet)] = true
                    }
                }
            }
        }
    }
    return ans
}
