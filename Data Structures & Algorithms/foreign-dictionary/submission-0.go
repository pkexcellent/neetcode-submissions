func foreignDictionary(words []string) string {
	adj := make(map[rune][]rune)
	indegree := make(map[rune]int)
	for _, w := range words {
		for _, c := range w {
			adj[c] = []rune{}
			indegree[c] = 0
		}
	}
   	for i := 0; i < len(words)-1; i++ {
		w1 := words[i]
		w2 := words[i+1]
		if len(w1) > len(w2) && w1[:len(w2)] == w2[:len(w2)] {
			return ""
		}
		for j := 0; j < len(w1) && j < len(w2); j++ {
			if w1[j] != w2[j] {
				r1, r2 := rune(w1[j]), rune(w2[j])
				adj[r1] = append(adj[r1], r2)
				indegree[r2]++
				break
			}
		}
   	} 

   	alphabet := ""
	q := []rune{}
	for c, in := range indegree {
		if in == 0 {
			q = append(q, c)
		}
	}
	for len(q) > 0 {
		cur := q[0]
		alphabet += string(cur)
		q = q[1:]
		for _, to := range adj[cur] {
			indegree[to]--
			if indegree[to] == 0 {
				q = append(q, to)
			}
		}
	}
	for _, in := range indegree {
		if in > 0 {
			return ""
		}
	}
	return alphabet
   	
}
