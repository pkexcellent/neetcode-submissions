func minSubArrayLen(target int, nums []int) int {
	// solution1 can only be used when nums are all positive
	// if there is any negative, we need to switch to prefixsum
    n := len(nums)
    prefix := make([]int, n+1)

    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }

    window := n + 1
    monotonicIdxQ := []int{0}

    for i := 1; i <= n; i++ {
        for len(monotonicIdxQ) > 0 && 
		   prefix[i] - prefix[monotonicIdxQ[0]] >= target {
            window = min(window, i - monotonicIdxQ[0])
            monotonicIdxQ = monotonicIdxQ[1:]
        }

		// maintain the monotonicIdxQ as a monotonic queue
		// prefix[i] > prefix[j], i > j
        for len(monotonicIdxQ) > 0 && 
		    prefix[i] <= prefix[monotonicIdxQ[len(monotonicIdxQ)-1]] {
            monotonicIdxQ = monotonicIdxQ[:len(monotonicIdxQ)-1]
        }

        monotonicIdxQ = append(monotonicIdxQ, i)
    }

    if window == n+1 {
        return 0
    }
    return window
}
