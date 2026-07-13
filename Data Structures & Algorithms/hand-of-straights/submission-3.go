func isNStraightHand(hand []int, groupSize int) bool {
    // loop the number starting with the smallest
    m := make(map[int]int)
    for _, num := range hand {
        m[num]++
    }
    sort.Ints(hand)
    for _, num := range hand {
        if m[num] == 0 {
            continue
        }
        start := num
        for i := start; i <= start + groupSize - 1; i++ {
            if v, exist := m[i]; !exist {
                return false
            } else {
                if v == 0 {
                    return false
                }
                m[i]--
            }
        }
    }
    return true
}
