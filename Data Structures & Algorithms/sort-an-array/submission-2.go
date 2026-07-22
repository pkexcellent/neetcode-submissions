func sortArray(nums []int) []int {
    //merge sort
	tmp := make([]int, len(nums))
	mergeSort(nums, tmp, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, tmp []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(nums, tmp, left, mid)
	mergeSort(nums, tmp, mid+1, right)
	merge(nums, tmp, left, mid, right)
}
func merge(nums []int, tmp []int, left, mid, right int) {
	copy(tmp[left:right+1], nums[left:right+1])
	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
		k++
	}
	for i <= mid {
		nums[k] = tmp[i]
		i++
		k++
	}
	for j <= right {
		nums[k] = tmp[j]
		j++
		k++
	}
}
