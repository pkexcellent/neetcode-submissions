func getSum(a int, b int) int {
	// use ^ without carry  <- sumNoCarry
	// use & << 1 to get carry <- carry
	// check if carry is 0, is 0 means stopping loop. 
	// then 

	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
		//fmt.Println(a, b, carry)
	}
	return a
	
}
