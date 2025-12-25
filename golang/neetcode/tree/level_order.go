package tree

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		var vs []int
		var nxt []*TreeNode

		for _, nd := range que {
			vs = append(vs, nd.Val)
			if nd.Left != nil {
				nxt = append(nxt, nd.Left)
			}
			if nd.Right != nil {
				nxt = append(nxt, nd.Right)
			}
		}
		ans = append(ans, vs)
		que = nxt
	}
	return
}
