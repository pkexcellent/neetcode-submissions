func longestConsecutive(nums []int) int {
    var m = make(map[int]bool, len(nums))
    for _, num := range nums {
        m[num] = false
    }
    var ans = 0
    for _, num := range nums {
        var i = 1
        if m[num] == true {
            continue
        }
        for {
            if _, exist := m[num+i]; exist {
                m[num+i] = true
                i++
            } else {
                break
            }
        }
        ans = max(ans, i)
    }
    return ans
}
