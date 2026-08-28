type Node struct {
	children map[rune]*Node
	endOfWord bool
}
type TrieTree struct {
	root *Node
}
func newTrieTree() *TrieTree {
	tt := &TrieTree {
		root: &Node {
			children: make(map[rune]*Node),
			endOfWord: false,
		},
	}
	return tt
}
func (tt *TrieTree) Insert(word string) {
	cur := tt.root
	for _, c := range word {
		if _, exist := cur.children[c]; !exist {
			cur.children[c] = &Node{children: make(map[rune]*Node)}
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true
}
func wordBreak(s string, wordDict []string) []string {
	// use a trie tree to lookup
	// use dfs to loop on the trie tree
	tt := newTrieTree()
	for _, w := range wordDict {
		tt.Insert(w)
	}
	// break s based on dfs
	rs := []string{}
	one := []string{}
	var dfs func(start, end int, node *Node)
	dfs = func(start, end int, node *Node) {
		if end == len(s) {
			if node == tt.root { // only when exact match a word and reach to an end
				oners := strings.Join(one, " ")
				rs = append(rs, oners)
			}
			return
		}
		if child, exist := node.children[rune(s[end])]; !exist {
			return
		} else {
			if child.endOfWord {
				phase := s[start:end+1]
				one = append(one, phase)
				dfs(end+1, end+1, tt.root)
				one = one[:len(one)-1]
				dfs(start, end+1, child)
			} else {
				dfs(start, end+1, child)
			}
		}
		return 
	}
	dfs(0, 0, tt.root)
	return rs
}
