func minEnd(n int, x int) int {
	// solution1 TLE
	// I can use n to build 1 to n-1, (x is 0)
	// and then insert x based on the 1s' position to generate the number
	largest := n-1
	insertBits := []int{}
	bit := 0
	for x > 0 {
		if x & 0x01 > 0 {
			insertBits = append(insertBits, bit)
		}
		bit++
		x = x >> 1
	}
	//fmt.Println(insertBits, largest)
	rs := 0
	bit = 0
	for len(insertBits) > 0 || largest > 0 {
		if len(insertBits) > 0 && bit == insertBits[0] {
			insertBits = insertBits[1:len(insertBits)]
			rs = rs + (1 << bit)
		} else {
			curbit := largest & 0x01
			rs = rs + (curbit << bit)
			largest = largest >> 1
		}
		bit++
		//fmt.Println(bit-1, largest, insertBits, rs)
	}
	return rs
}
