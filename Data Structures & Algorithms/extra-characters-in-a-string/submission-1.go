// build trie tree
type TrieNode struct {
	children map[rune]*TrieNode
	endOfWord bool
}

type TrieTree struct {
	root *TrieNode
}
func NewTrieTree() *TrieTree {
	return &TrieTree{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}
func(t *TrieTree) Insert(word string) {
	node := t.root
	for _, c := range word {
		if child, exist := node.children[c]; exist {
			node = child
		} else {
			newChild := &TrieNode {
				children: make(map[rune]*TrieNode),
			}
			node.children[c] = newChild
			node = newChild
		}
	}
	node.endOfWord = true
}

func minExtraChar(s string, dictionary []string) int {
	// trie tree + dfs
	trieTree := NewTrieTree()
	for _, word := range dictionary {
		trieTree.Insert(word)
	}

	//fmt.Println(*trieTree.root)
	// dfs string based on trieTree
	endIdx := len(s)-1
	memo := make(map[int]int)
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx == len(s) {
			return 0
		}
		if rs, exist := memo[idx]; exist {
			return rs
		}
		best := 1 + dfs(idx + 1)
		node := trieTree.root
		for i := idx; i <= endIdx; i++ {
			cur := rune(s[i])
			if next, exist := node.children[cur]; exist {
				if next.endOfWord {
					best = min(best, dfs(i+1))
				}
				node = next
			} else {
				break
			}
		}
		memo[idx] = best
		return best
	}
	return dfs(0)
}
