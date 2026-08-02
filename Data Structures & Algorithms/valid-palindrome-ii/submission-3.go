func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	cnt := 0
	for l < r {
		fmt.Println(string(s[l]), string(s[r]), l, r)
		if s[l] != s[r] {
			cnt++
			if cnt > 1 {
				return false
			}
			return valid(s, l+1, r) || valid(s, l, r-1)
		} else {
			l++
			r--
		}
	}
	return true
}

func valid(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
