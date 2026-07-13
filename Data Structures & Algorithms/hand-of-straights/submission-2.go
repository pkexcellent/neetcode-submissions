func isNStraightHand(hand []int, groupSize int) bool {
    // re-order the list, and start from the smallest one
    // pick groupSize numbers +1
    //  find the first repeated number, which will be the smallest 
    // repeatly
    // build a map is better
    handMap := make(map[int]int)
    startNumber := math.MaxInt32
    for _, num := range hand {
        handMap[num]++
        if num < startNumber {
            startNumber = num
        }
    }
    //fmt.Println(handMap, startNumber)

    var dfs func(handMap map[int]int, startNumber int) bool
    dfs = func(handMap map[int]int, startNumber int) bool {
        if len(handMap) == 0 {
            return true
        }
        if len(handMap) < groupSize {return false}
        for i := startNumber; i <= startNumber+groupSize-1; i++ {
            if v, exist := handMap[i]; !exist {
                return false
            } else {
                if v == 0 {
                    return false
                }
                handMap[i]--
                if handMap[i] == 0 {
                    delete(handMap, i)
                }
            }
        }
        newStart := math.MaxInt32
        for k, _ := range handMap {
            if k < newStart {
                newStart = k
            }
        }
        //fmt.Println(handMap, newStart)
        return dfs(handMap, newStart)   
    }
    return dfs(handMap, startNumber)
}
