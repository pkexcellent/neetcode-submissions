func canCompleteCircuit(gas []int, cost []int) int {
    // this is like the sub array with always positive sum
    n := len(gas)
    remains := make([]int, n)
    for i := 0; i < n; i++ {
        remains[i] = gas[i] - cost[i]
    }
    // now remains means the each delta for stations, 
    // we need to start with a positive one, and while adding, the sum should always be positive
    startPos := -1
    for i := 0; i < n; i++ {
        if remains[i] < 0 {
            continue
        }
        streakSum := 0 // start loop
        canDo := true
        for j := i; j < n; j++ {
            streakSum += remains[j]
            if streakSum < 0 {
                canDo = false
                break
            }
        }
        if !canDo {
            continue
        }
        for j := 0; j < i; j++ {
            streakSum += remains[j]
            if streakSum < 0 {
                canDo = false
                break
            }
        }
        if canDo {
            startPos = i
            break
        }  
    }
    return startPos
}
