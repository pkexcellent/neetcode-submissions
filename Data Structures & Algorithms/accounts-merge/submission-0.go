type UnionSet struct {
	parent []int
	rank []int
}
func NewUnionSet(n int) *UnionSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i, _ := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionSet {
		parent,
		rank,
	}
}
func (us *UnionSet) Find(x int) int {
	rs := us.parent[x]
	if x != rs {
		rs = us.Find(rs)
	}
	return rs
}
func (us *UnionSet) Merge(x, y int) int {
	px, py := us.Find(x), us.Find(y)
	if px == py {
		return px
	}
	if us.rank[px] >= us.rank[py] {
		us.rank[px] += us.rank[py]
		us.parent[py] = px
		return px
	} else {
		us.rank[py] += us.rank[px]
		us.parent[px] = py
		return py
	}
}

func accountsMerge(accounts [][]string) [][]string {
	// build a map with email:name_idx
	// same email will lead to name_idx merge
	// the general idea is correct, but for judging how the accounts are linked
	// we need to use union set
	n := len(accounts)
	us := NewUnionSet(n)
	accountM := make(map[string][]int)
	for idx, account := range accounts {
		for i := 1; i < len(account); i++ {
			accountM[account[i]] = append(accountM[account[i]], idx)
			if len(accountM[account[i]]) > 1 {
				parent := us.Merge(idx, accountM[account[i]][0])
				accountM[account[i]] = []int{parent}
			}
		}
	}
	fmt.Println(accountM)
	fmt.Println(us)

	rs := make(map[int][]string)
	for email, parent := range accountM {
		rootIdx := us.Find(parent[0])
		rs[rootIdx] = append(rs[rootIdx], email)
	}
	stringRs := [][]string{}
	for idx, accs := range rs {
		name := accounts[idx][0]
		stringRs = append(stringRs, append([]string{name}, accs...))
	}
	return stringRs

}
