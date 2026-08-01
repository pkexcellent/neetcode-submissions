func subarraySum(nums []int, k int) int {
	// use prefix sum
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	cnt := 0
	totalSum := 0
	for _, num := range nums {
		totalSum += num
		// should be totalSum -k, because we use prefix, so the sub array sum is 
		// prefix[i] - prefix[j] = sum - (sum-k) = k
		if v, exist := prefixSum[totalSum-k]; exist {
			cnt += v
		}
		prefixSum[totalSum]++
	}
	return cnt
}
