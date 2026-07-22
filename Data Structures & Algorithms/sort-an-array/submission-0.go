func sortArray(nums []int) []int {
    // quick sort
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

func partition(nums []int, left, right int) int {
	i := left - 1
	pivot := nums[right]
	for j := left; j <= right; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
