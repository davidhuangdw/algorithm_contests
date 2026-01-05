package tree

func preorderTraversal1(root *TreeNode) []int { // by iter
	var nxt []*TreeNode
	cur, ans := root, make([]int, 0)
	for !(cur == nil && len(nxt) == 0) {
		for cur != nil {
			ans = append(ans, cur.Val)
			nxt = append(nxt, cur.Right)
			cur = cur.Left
		}
		cur, nxt = nxt[len(nxt)-1], nxt[:len(nxt)-1]
	}
	return ans
}

func preorderTraversal(root *TreeNode) []int { // by morris
	cur, ans := root, make([]int, 0)
	for cur != nil {
		if cur.Left == nil {
			ans = append(ans, cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != cur && pre.Right != nil {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur // attach
			ans = append(ans, cur.Val)
			cur = cur.Left
		} else { // left done
			pre.Right = nil // detach
			cur = cur.Right
		}
	}
	return ans
}
