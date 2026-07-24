func sortColors(nums []int) {
    r, w, b := 0, 0, 0
	for _, n := range nums {
		if n == 0 {
			r++
		} else if n == 1 {
			w++
		} else {
			b++
		}
	}
	for i := 0; i < r; i++ {
		nums[i] = 0
	}
	for i := r; i < r+w; i++ {
		nums[i] = 1
	}
	for i := r+w; i < r+w+b; i++ {
		nums[i] = 2
	}
}
