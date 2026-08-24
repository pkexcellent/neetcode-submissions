func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n == 0 {return 0}
	l, r := 0, 0
	subsum := 1
	longest := subsum
	preDirection := 0
	for i := 1; i < n; i++ {
		direction := 0
		if arr[i] > arr[i-1] {
			direction = 1
		} else if arr[i] < arr[i-1] {
			direction = -1
		} else {
			direction = 0
		}
		if direction == 0 { // same number, reset l as current
			l = i
			longest = max(longest, subsum)
		} else if preDirection != direction {
			r = i
			subsum = r - l + 1
		} else {
			l = i-1
		}
		preDirection = direction
	}
	longest = max(longest, subsum)
	return longest
}
