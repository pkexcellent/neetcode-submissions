func addBinary(a string, b string) string {
	rs := []rune{}
	ar, br := reverse([]rune(a)), reverse([]rune(b))

	na, nb := len(ar), len(br)
	if na < nb {
		ar, br = br, ar
		na, nb = nb, na
	}
	carry := 0
	for i := 0; i < na; i++ {
		sum := int(ar[i] - '0') + carry
		if i < nb {
			sum += int(br[i] - '0')
		}
		carry = sum / 2
		rs = append(rs, '0' + rune(sum%2))
	}
	if carry != 0 {
		rs = append(rs, '0' + rune(carry))
	}
	return string(reverse(rs))
}
func reverse(a []rune) []rune {
	n := len(a)-1
	for i := 0; i <= n/2; i++ {
		a[i], a[n-i] = a[n-i], a[i]
	}
	return a
}
