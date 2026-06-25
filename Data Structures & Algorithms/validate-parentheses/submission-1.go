func isValid(s string) bool {
    var stack = []rune{}
	for _, c := range []rune(s) {
		if c == '(' ||  c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			lastC := stack[len(stack)-1]
			if (c == ')' && lastC == '(') || 
			   (c == '}' && lastC == '{') ||
			   (c == ']' && lastC == '[') {
				stack = stack[0:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
