func openLock(deadends []string, target string) int {
	dead := make(map[string]struct{})
	for _, deadend := range deadends {
		dead[deadend] = struct{}{}
	}
	if _, exist := dead["0000"]; exist {
		return -1
	}
	path := []string{"0000"}
	dead["0000"] = struct{}{}

	step := 0
	for len(path) > 0 {
		cases := len(path)
		step++
		for i := 0; i < cases; i++ {
			start := path[0]
			path = path[1:]
			for _, delta := range []int{1, -1} {
				for j := 0; j < 4; j++ {
					digChange := rune((int(start[j] - '0') + delta + 10) % 10) + '0'
					newcom := start[0:j] + string(digChange) + start[j+1:]
					if newcom == target {
						return step
					}
					if _, exist := dead[newcom]; exist {
						continue
					}
					path = append(path, newcom)
					dead[newcom] = struct{}{}
				}
			}
		}
	}
	return -1
}
