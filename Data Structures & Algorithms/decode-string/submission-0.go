func decodeString(s string) string {
	// easy
	stack := []string{}
	for _, c := range s {
		if string(c) == "]" {
			subs := ""
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				popc := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				subs = string(popc) + subs
			}
			stack = stack[:len(stack)-1] // pop "["
			repeatn := ""
			for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				repeatn = stack[len(stack)-1] + repeatn
				stack = stack[:len(stack)-1]
			}
			repeat, _ := strconv.Atoi(repeatn)
			//fmt.Println("repeat", repeat)
			pushs := ""
			for repeat > 0 {
				pushs += subs
				repeat--
			}
			stack = append(stack, pushs)
		} else {
			stack = append(stack, string(c))
		}
	}
	//fmt.Println(stack)
	rs := ""
	for len(stack) > 0 {
		rs = stack[len(stack)-1] + rs
		stack = stack[:len(stack)-1]
	}
	return rs
}
