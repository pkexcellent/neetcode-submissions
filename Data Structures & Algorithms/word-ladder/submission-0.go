func ladderLength(beginWord string, endWord string, wordList []string) int {
    // use BFS-like method to find the shorest path
	
	// build a worklist map to be easily lookup
	words := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		words[word] = true
	}

	// use BFS to find layer by layer, with 1-distinct word, if exists
	q := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	layer := 1
	for len(q) > 0 {
		layer++
		size := len(q)
		for idx := 0; idx < size; idx++ {
			cur := q[0]
			q = q[1:]
			for c := 'a'; c <= 'z'; c++ {
				for i := 0; i < len(cur); i++ {
					tmp := cur[:i] + string(c) + cur[i+1:]
					if visited[tmp] || !words[tmp] {
						continue
					}
					// if found, the distince is 1
					if tmp == endWord {
						return layer
					}
					q = append(q, tmp)
					visited[tmp] = true
				} 
			}
		}
	}
	return 0

}
