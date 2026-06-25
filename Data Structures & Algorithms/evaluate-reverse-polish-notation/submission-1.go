func evalRPN(tokens []string) int {
	var stack = []int{}
	for _, c := range tokens {
		if c != "+" && c != "-" && c != "*" && c != "/" {
			num, _ := strconv.Atoi(c)
			stack = append(stack, num)
		} else {
			lastNum := stack[len(stack)-2]
			num := stack[len(stack)-1]
			if c == "+" {
				stack[len(stack)-2] = lastNum + num
			} else if c == "-" {
				stack[len(stack)-2] = lastNum - num
			} else if c == "*" {
				stack[len(stack)-2] = lastNum * num
			} else if c == "/" {
				stack[len(stack)-2] = lastNum / num
			}
			stack = stack[:len(stack)-1]
		}
	}
	return stack[0]
}
