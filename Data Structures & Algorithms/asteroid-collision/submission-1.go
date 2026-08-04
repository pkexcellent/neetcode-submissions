func asteroidCollision(asteroids []int) []int {
	// find adjencent elements, merge them
	// repeat, until no elements can be merged
	stack := []int{}
	for _, ast := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, ast)
		} else {
			top := stack[len(stack)-1]
			newAst := ast
			for newAst < 0 && top > 0 && len(stack) > 0 {
				stack = stack[:len(stack)-1]
				if abs(newAst) > top {
					newAst = newAst
					if len(stack) > 0 {
						top = stack[len(stack)-1]
					}
				} else if abs(newAst) < top {
					newAst = top
				} else {
					newAst = 0
				}
			}
			if newAst != 0 {
				stack = append(stack, newAst)
			}
		}
	}
	return stack
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
