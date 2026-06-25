// trie tree
type TrieNode struct {
	children map[rune]*TrieNode
	endOfWord bool
}
func constructTrie() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

func (t *TrieNode) Insert (word string) {
	cur := t
	for _, c := range word {
		if _, exist := cur.children[c]; !exist {
			cur.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true
}

func findWords(board [][]byte, words []string) []string {
    t := constructTrie()
	for _, word := range words {
		t.Insert(word)
	}

	var visited = make([][]bool, len(board))
	for i, _ := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	var found = make(map[string]struct{})

	var dfs func(r, c int, t *TrieNode, path string)
	dfs = func(r, c int, t *TrieNode, path string) {
		if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || visited[r][c] {
			return
		}
		char := rune(board[r][c])
		if _, exist := t.children[char]; exist {
			visited[r][c] = true
			path += string(char)
			// note! this is to judge the children endOfWord, not parent
			if t.children[char].endOfWord == true {
				found[path] = struct{}{}
			}
			dfs(r+1, c, t.children[char], path)
			dfs(r-1, c, t.children[char], path)
			dfs(r, c+1, t.children[char], path)
			dfs(r, c-1, t.children[char], path)
			visited[r][c] = false
		} else {
			return
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(i, j, t, "")
		}
	}
	rs := []string{}
	for k, _ := range found {
		rs = append(rs, k)
	}
	return rs
}
