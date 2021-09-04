package leetcode



func preorderTraversal(root *TreeNode) []int {
	var st []*TreeNode
	var pop = func() *TreeNode{
		l := len(st)
		v := st[l-1]
		st = st[:l-1]
		return v
	}

	x := root
	var order []int
	for x != nil || len(st) > 0{
		if x != nil {
			order = append(order, x.Val)
			if x.Right != nil{
				st = append(st, x.Right)
			}
			x = x.Left
		}else{
			x = pop()
		}
	}
	return order
}
