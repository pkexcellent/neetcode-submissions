func permute(nums []int) [][]int {
	var rs = [][]int{}
	var cur = []int{}
	rs = append(rs, cur)
	for _, num := range nums {
		newRs := [][]int{}
		for _, com := range rs {
			for i := 0; i <= len(com); i++ {
				tmp := append([]int{}, com...)
				tmp = append(tmp[0:i], append([]int{num}, tmp[i: len(tmp)]...)...)
				newRs = append(newRs, tmp)
			}
		}
		rs = newRs
	}

	return rs
}
