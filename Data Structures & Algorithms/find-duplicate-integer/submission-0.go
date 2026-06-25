func findDuplicate(nums []int) int {
	for _, num := range nums {
		if nums[abs(num)-1] < 0 {
			return abs(num)
		} else {
			nums[abs(num)-1] = nums[abs(num)-1]*(-1)
		}
	}
	return 0
}

func abs(num int) int {
	if num < 0 {
		return num*(-1)
	}
	return num
}
