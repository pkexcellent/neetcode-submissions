func canCompleteCircuit(gas []int, cost []int) int {
    // sliding window
    n := len(gas)
    if n < 2 {
        if (gas[0]-cost[0]) >= 0 {
            return 0
        } else {
            return -1
        }
    }
    start, end := 0, 1
    tank := gas[start]-cost[start]
    for start != end {
        if tank < 0 {
            start = (start-1 + n)%n
            tank += gas[start]-cost[start]
        } else {
            tank += gas[end] - cost[end]
            end = (end+1 + n)%n // this should be after modify tank
        }
    }
    if tank >= 0 {
        return start
    }
    return -1
}
