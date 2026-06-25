type TrieNode struct {
	child map[rune]*TrieNode
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: &TrieNode{
			child: make(map[rune]*TrieNode),
			endOfWord: true,
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	trie := this.root
	/*
	for _, c := range word {
		if trie.child[c] == nil {
			trie.child[c] = &TrieNode{child: make(map[rune]*TrieNode)}
		}
		trie = trie.child[c]
	}*/
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

func (this *PrefixTree) Search(word string) bool {
	trie := this.root
	for _, c := range word {
		if _, exist := trie.child[c]; exist {
			trie = trie.child[c]
		} else {
			return false
		}
	}
	return trie.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	trie := this.root
	for _, c := range prefix {
		if _, exist := trie.child[c]; exist {
			trie = trie.child[c]
		} else {
			return false
		}
	}
	return true
}
