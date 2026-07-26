func majorityElement(nums []int) int {
	if len(nums) == 0 {return -1 }// >??
    cur := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			cur = nums[i]
			cnt++
		} else if nums[i] == cur {
			cnt++
		} else {
			cnt--
		}
	}
	return cur
}
