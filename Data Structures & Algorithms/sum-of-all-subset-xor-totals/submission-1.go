func subsetXORSum(nums []int) int {
	var dfs func(idx int, total int) int
	dfs = func(idx int, total int) int {
		if idx == len(nums) {
			return total
		}
		return dfs(idx+1, total^nums[idx]) + dfs(idx+1, total)
	}
	return dfs(0, 0)
}
