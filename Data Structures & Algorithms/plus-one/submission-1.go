func plusOne(digits []int) []int {
    carry := 0
	n := len(digits)
	if digits[n-1] < 9 {
		digits[n-1]++
		return digits
	}
	for i := n-1; i >= 0; i-- {	
		newNum := digits[i] + carry
		if i == n - 1 {
			newNum += 1 
		}
		carry = newNum / 10
		newNum = newNum % 10
		digits[i] = newNum
		if carry == 0 {
			break
		}
	}
	if carry > 0 {
		return append([]int{carry}, digits...)
	} else {
		return digits
	}
}
