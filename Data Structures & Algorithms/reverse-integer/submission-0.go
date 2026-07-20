func reverse(x int) int {
	newNum := 0
	sign := 1
	if x < 0 {
		sign = -1
		x = -1 * x
	}
	fmt.Println(x)
	for x > 0 {
		lastDigit := x % 10
		if newNum > 0x7FFFFFFF / 10 || 
			(newNum == 0x7FFFFFFF / 10 && lastDigit > 0x7FFFFFFF % 10){
			return 0
		}
		newNum = newNum*10 + lastDigit
		x = x / 10
	}
	return sign * newNum

	
}
