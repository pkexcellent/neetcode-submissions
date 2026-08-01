func firstMissingPositive(nums []int) int {
	// use existing position's sign to indicate if i exist
	// firstly we need to convert all negative numbers to 0 to avoid miss
	for i, _ := range nums {
		if nums[i] < 0 {
			nums[i] = 0
		}
	}
	for i, _ := range nums {
		if nums[i] <= 0 {
			continue
		}
		swapIdx := nums[i]-1
		for nums[i] != i+1 && 
			swapIdx >= 0 && swapIdx < len(nums) && 
			nums[i] != nums[swapIdx] {
			nums[i], nums[swapIdx] = nums[swapIdx], nums[i]
			swapIdx = nums[i]-1
		}
	}
	for i, num := range nums {
		if num == 0 || num != i+1 {
			return i+1
		}
	}
	return len(nums)+1
}
