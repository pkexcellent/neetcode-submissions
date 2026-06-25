func hasDuplicate(nums []int) bool {
    var numsMap = make(map[int]bool)
    for _, num := range nums {
        if _, exist := numsMap[num]; exist {
            return true
        } else {
            numsMap[num] = true
        }
    }
    return false
}
