type TrieNode struct {
	child map[rune]*TrieNode
	wordEnd bool
}
type TrieNodeTree struct {
	root *TrieNode
}

func (t *TrieNodeTree) Insert(word string) {
	cur := t.root
	for i := len(word)-1; i >= 0; i-- {
		c := rune(word[i])
		if _, exist := cur.child[c]; !exist {
			cur.child[c] = &TrieNode{child: make(map[rune]*TrieNode)} 
		}
		cur = cur.child[c]
	}
	cur.wordEnd = true
}
func (t *TrieNodeTree) Get(word string) bool {
	cur := t.root
	for _, c := range word {
		if _, exist := cur.child[c]; !exist {
			return false
		}
		cur = cur.child[c]
	}
	return cur.wordEnd
}

func minExtraChar(s string, dictionary []string) int {
	// build a trie tree
	// use DP
	// dp[i] is from 0-i, how many extra chars
	// dp[i+1] = dp[i] + 1 (if we treat this 1 as extrac), or
	// j, (j< i), dp[j] + s[j+1:i] , min of all j
	// use dictionary to build tree

	tt := &TrieNodeTree {root: &TrieNode{child: make(map[rune]*TrieNode)}}
	// has to insert with a reversed word
	for _, word := range dictionary {
		tt.Insert(word)
	}

	dp := make([]int, len(s)+1)
	dp[0] = 0
	for i := 1; i < len(s) + 1; i++ {
		dp[i] = dp[i-1] + 1 // potential max dp[i]
		cur := tt.root
		for j := i-1; j >=0; j-- {
			if _, exist := cur.child[rune(s[j])]; !exist {
				break
			}
			cur = cur.child[rune(s[j])]
			if cur.wordEnd {
				dp[i] = min(dp[j], dp[i])
			}
		}
	}
	return dp[len(s)]
}
