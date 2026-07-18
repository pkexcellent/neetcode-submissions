func plusOne(digits []int) []int {
    b := 0
	n := len(digits)
	for i := n-1; i >= 0; i-- {	
		newNum := digits[i] + b
		if i == n - 1 {
			newNum += 1 
		}
		b = newNum / 10
		newNum = newNum % 10
		digits[i] = newNum
	}
	if b > 0 {
		return append([]int{b}, digits...)
	} else {
		return digits
	}
}
