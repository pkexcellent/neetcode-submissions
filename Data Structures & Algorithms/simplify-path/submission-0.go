func simplifyPath(path string) string {
	// use a stack to simulate the path
	stack := []string{}
	eles := strings.Split(path, "/")
	//fmt.Println(eles)
	for _, ele := range eles {
		if ele == "/" || ele == "." || ele == "" {
			continue
		} 
		if ele == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, "/" + ele)
		}
	}
	if len(stack) == 0 {
		return "/"
	} else {
		return strings.Join(stack, "")
	}
}
