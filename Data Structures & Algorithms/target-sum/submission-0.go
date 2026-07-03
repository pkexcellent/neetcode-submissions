func findTargetSumWays(nums []int, target int) int {
    // use dfs, each number can be used as minus or positive
	count := 0
	var dfs func(sum int, idx int) int
	dfs = func(sum, idx int) int {
		if idx > len(nums) {
			return 0
		}
		if idx == len(nums) {
			if sum == target {
				count++
				return 1
			} else {
				return 0
			}
		}


		// use this number as positive
		dfs(sum+nums[idx], idx+1)
		// or use this as negetive
		dfs(sum-nums[idx], idx+1)
		return 0
	}
	dfs(0, 0)
	return count
}
