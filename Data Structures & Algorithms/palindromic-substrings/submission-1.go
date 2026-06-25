func countSubstrings(s string) int {
    count := 0 
	n := len(s)
	for i := 0; i < n; i++ {
		for del := 0; del < n; del++ {
			l, r := i - del, i + del
			if l < 0 || r >= n || s[l] != s[r] {
				break
			}
			count++
			//fmt.Println(s[l:r+1])
		}
	} 

	for i := 0; i < n; i++ {
		j := i+1
		if j >= n || s[j] != s[i] {
			continue
		}
		for del := 0; del < n; del ++ {
			l, r := i-del, j+del
			if l < 0 || r >= n || s[l] != s[r] {
				break
			}
			count++
			//fmt.Println(s[l:r+1])
		}
	}
	return count
}
