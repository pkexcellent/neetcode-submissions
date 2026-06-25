/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    var ans = []int{}
	if root == nil {return ans}
	var qu = []*TreeNode{root}
	for len(qu) > 0 {
		l := len(qu)
		thisLevel := []int{}
		for i := 0; i<l; i++ {
			thisLevel = append(thisLevel, qu[0].Val)
			if qu[0].Left != nil {
				qu = append(qu, qu[0].Left)
			}
			if qu[0].Right != nil {
				qu = append(qu, qu[0].Right)
			}
			qu = qu[1:]
		}
		//fmt.Println(thisLevel)
		ans = append(ans, thisLevel[len(thisLevel)-1])
	}
	return ans
}
