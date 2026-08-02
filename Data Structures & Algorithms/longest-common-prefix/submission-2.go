type TrieNode struct {
	child map[rune]*TrieNode
	endOfWord bool
}

type TrieTree struct {
	root *TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: &TrieNode {
			child: make(map[rune]*TrieNode),
			endOfWord: false,
		},
	}
}
func (t *TrieTree) Insert(word string) {
	cur := t.root
	for _, c := range word {
		if _, exist := cur.child[c]; exist {
			cur = cur.child[c]
			continue
		}
		cur.child[c] = &TrieNode{
			child: make(map[rune]*TrieNode),
			endOfWord: false,
		}
		cur = cur.child[c]
	}
	cur.endOfWord = true
}

func (t *TrieTree) findPrefixLength(word string) int {
	cur := t.root
	for i, c := range word {
		if _, exist := cur.child[c]; exist {
			cur = cur.child[c]
		} else {
			return i
		}
	}
	return len(word)
}

func longestCommonPrefix(strs []string) string {
    // trie tree
	// find the shortest str
	n := len(strs)
	if n == 0 {
		return ""
	}
	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}
	tt := NewTrieTree()
	tt.Insert(shortest)
	l := len(shortest)
	for _, str := range strs {
		l = min(l, tt.findPrefixLength(str))
	}
	return shortest[:l]
	
}
