func reverseBits(n int) int {
	reversed := 0
	for i := 0; i < 32; i++ {
		lowbit := n & 1
		//fmt.Println(fmt.Sprintf("%032b", lowbit))
		reversed = (reversed << 1) | lowbit
		//fmt.Println(fmt.Sprintf("%032b", reversed))
		n = n >> 1
	}
	//fmt.Println(fmt.Sprintf("%032b", reversed))
	return reversed
}
