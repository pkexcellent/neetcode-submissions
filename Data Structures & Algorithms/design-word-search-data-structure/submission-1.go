type TrieNode struct {
	child map[rune]*TrieNode
	endOfWord bool
}

type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
		root: &TrieNode {
			child: make(map[rune]*TrieNode),
		},
	}
}

func (this *WordDictionary) AddWord(word string)  {
	trie := this.root
	for _, c := range word {
		if _, exist := trie.child[c]; exist {
			trie = trie.child[c]
		} else {
			trie.child[c] = &TrieNode{child: make(map[rune]*TrieNode)}
			trie = trie.child[c]
		}
	}
	trie.endOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
    trie := this.root
	var dfs func(idx int, trie *TrieNode) bool
	dfs = func(idx int, trie *TrieNode) bool {
		for i := idx; i < len(word); i++ {
			c := rune(word[i])
			if c != rune('.') {
				if _, exist := trie.child[c]; !exist {
					return false
				}
				trie = trie.child[c]
			} else {
				for _, child := range trie.child {
					if child != nil && dfs(i+1, child) {
						return true
					}
				}
				return false
			}
		}
		return trie.endOfWord
	}
	return dfs(0, trie)
}
