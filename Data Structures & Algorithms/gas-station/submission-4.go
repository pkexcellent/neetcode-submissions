func canCompleteCircuit(gas []int, cost []int) int {
    // this is the real method like max sub array
    // once the sum is negative, we should reset it as zero
    sumGas := 0
    for i := 0; i < len(gas); i ++ {
        sumGas += gas[i] - cost[i]
    }
    if sumGas < 0 {
        return -1
    }

    curSum := 0
    startPos := 0
    for i := 0; i < len(gas); i++ {
        curSum += gas[i] - cost[i]
        if curSum < 0 {
            curSum = 0
            startPos = i+1
        }
    }
    return startPos
}
