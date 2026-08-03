func calPoints(operations []string) int {
	stack := []int{}
	for _, operation := range operations {
		if operation == "+" {
			n := len(stack)-1
			num1, num2 := stack[n-1], stack[n]
			stack = append(stack, num1+num2)
		} else if operation == "D" {
			n := len(stack)-1
			stack = append(stack, stack[n]*2)
		} else if operation == "C" {
			stack = stack[:len(stack)-1]
		} else {
			num, _ := strconv.Atoi(operation)
			stack = append(stack, num)
		}
	}
	sum := 0
	for _, n := range stack {
		sum += n
	}
	return sum
}
