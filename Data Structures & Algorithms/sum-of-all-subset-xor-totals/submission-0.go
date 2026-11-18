func subsetXORSum(nums []int) int {
	// get all the subset
	subsets := [][]int{}
	var backtrack func(idx int, subset []int) 
	backtrack = func(idx int, subset []int) {
		tmp := make([]int, len(subset))
		copy(tmp, subset)
		subsets = append(subsets, tmp)

		for i := idx; i < len(nums); i++ {
			subset = append(subset, nums[i])
			backtrack(i+1, subset)
			subset = subset[:len(subset)-1]
		}
	}
	backtrack(0, []int{})
	//fmt.Println(subsets)
	rs := 0
	for _, sb := range subsets {
		localRs := 0
		for _, num := range sb {
			localRs ^= num
		}
		rs += localRs
	}
	return rs
}
