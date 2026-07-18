func isHappy(n int) bool {
    appeared := make(map[int]struct{})
	num := 0
	for {
		for n != 0 {
			dig := n % 10
			num = num + dig*dig
			n = n/10
		}
		if _, exist := appeared[num]; exist {
			return false
		}
		appeared[num] = struct{}{}
		if num == 1 {
			return true
		} else {
			n = num
			num = 0
		}
	}
	return true
}
