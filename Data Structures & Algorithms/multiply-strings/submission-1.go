func multiply(num1 string, num2 string) string {
    // wawawa
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	finalRs := ""
	for i := len(num2)-1; i >= 0; i-- {
		thisNum := int(num2[i] - '0')
		thisRs := productOf(num1, thisNum, len(num2)-1-i)
		finalRs = stringAdd(finalRs, thisRs)
	}
	return finalRs
}

func productOf(a string, b int, numOfZeros int) string {
	n := len(a)
	carry := 0
	rs := []byte{}
	for i := n-1; i >= 0; i-- {
		num := int(a[i] - '0')
		product := num * b + carry
		rs = append(rs, byte(product%10 + '0'))
		carry = product/10
	}
	if carry != 0 {
		rs = append(rs, byte(carry + '0'))
	}
	// reverse the array
	for i := 0; i < len(rs)/2; i++ {
		j := len(rs) - 1 - i
		rs[i], rs[j] = rs[j], rs[i]
	}
	result := string(rs)
	for i := 0; i < numOfZeros; i++ {
		result += "0"
	}
	return result
}

func stringAdd(a, b string) string {
	m, n := len(a)-1, len(b)-1
	rs := []byte{}
	carry := 0
	for m >= 0 || n >= 0 || carry > 0 {
		thisA, thisB := 0, 0
		if m >= 0 {
			thisA = int(a[m] - '0')
		}
		if n >= 0 {
			thisB = int(b[n] - '0')
		}
		add := thisA + thisB + carry
		carry = add/10
		rs = append([]byte{byte(add%10 + '0')}, rs...)
		m--
		n--
	}
	return string(rs)
}
