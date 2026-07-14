func checkValidString(s string) bool {
    // use stack, value is index
    stackL := []int{}
    stackStar := []int{}
    for i, c := range s {
        if c == '(' {
            stackL = append(stackL, i)
        } else if c == '*' {
            stackStar = append(stackStar, i)
        } else {
            if len(stackL) > 0 {
                stackL = stackL[:len(stackL)-1]
            } else if len(stackStar) > 0 {
                stackStar = stackStar[:len(stackStar)-1]
            } else {
                return false
            }
        }
    }
    for len(stackL) > 0 {
        lastL := stackL[len(stackL)-1]
        if len(stackStar) <= 0 {
            return false
        }
        lastStar := stackStar[len(stackStar)-1]
        if lastStar < lastL {
            return false
        }
        stackL = stackL[:len(stackL)-1]
        stackStar = stackStar[:len(stackStar)-1]
    }
    return true
}
