func subsetsWithDup(nums []int) [][]int {
	// still backtracking method
	var rs = [][]int{}
	sort.Ints(nums)

	var backtrack func(int, []int)
	backtrack = func(i int, tmp []int) {
		if i == len(nums) {
			tmpp := make([]int, len(tmp))
			copy(tmpp, tmp)
			rs = append(rs, tmpp)
			return
		}
		tmp = append(tmp, nums[i])
		backtrack(i+1, tmp)
		// not put nums[i]
		tmp = tmp[:len(tmp)-1]
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
		backtrack(i+1, tmp)
	}
	backtrack(0, []int{})
	return rs
}
