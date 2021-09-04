package leetcode

func inorderTraversal(root *TreeNode) []int {
	var order []int
	var st []*TreeNode
	pop := func() *TreeNode{
		l := len(st)
		v := st[l-1]
		st = st[:l-1]
		return v
	}

	x := root
	for x != nil || len(st) > 0 {
		if x != nil {
			st = append(st, x)
			x = x.Left
		}else{
			x = pop()
			order = append(order, x.Val)
			x = x.Right
		}
	}

	return order
}
