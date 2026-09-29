func mergeAlternately(word1 string, word2 string) string {
	idx1, idx2 := 0, 0
	n1, n2 := len(word1), len(word2)
	rs := ""
	for idx1 < n1 || idx2 < n2 {
		if idx1 < n1 {
			rs += string(word1[idx1])
			idx1++
		}
		if idx2 < n2 {
			rs += string(word2[idx2])
			idx2++
		}	
	}
	return rs
}
