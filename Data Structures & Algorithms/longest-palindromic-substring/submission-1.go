func longestPalindrome(s string) string {
    // if it is an odd, from the mid to sides
	size := 0
	maxS := ""

	// odd
	for i := 0; i < len(s); i++ {
		for delta := 0; delta < len(s); delta ++ {
			left := i - delta
			right := i + delta
			if left < 0 || right >= len(s) {
				break
			}
			if s[left] == s[right] {
				if delta*2 + 1 > size {
					size = delta*2 + 1
					maxS = s[left:right+1]
				}
			} else {
				break
			}
		}
	}

	// even
	for i := 0; i < len(s)-1; i++ {
		j := i+1
		if s[i] != s[j] {
			continue
		}
		for delta := 0; delta < len(s); delta++ {
			left := i-delta
			right := j + delta
			//fmt.Println(delta, left, right)
			if left < 0 || right >= len(s) || s[left] != s[right]{
				break
			}
			//fmt.Println(":::", size, delta*2 + 2, s[left:right+1])
			if delta*2 + 2 > size {
				size = delta*2 + 2
				maxS = s[left:right+1]
			}
		}
	}
	return maxS
}
