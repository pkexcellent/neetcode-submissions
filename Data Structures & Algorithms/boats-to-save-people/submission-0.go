func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l, r := 0, len(people)-1
	cnt := 0
	for l <= r {
		if r == l {
			cnt++
			l++
		} else if people[l] + people[r] > limit {
			cnt++
			r--
		} else if people[l] + people[r] == limit {
			cnt++
			l++
			r--
		} else {
			if people[r] + people[r-1] <= limit {
				r -= 2
				cnt++
			} else {
				cnt++
				r--
				l++
			}
		}
		//fmt.Println(l, r, cnt)
	}
	return cnt
}
