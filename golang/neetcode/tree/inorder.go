package tree

func inorderTraversal1(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(nd *TreeNode)
	dfs = func(nd *TreeNode) {
		if nd == nil {
			return
		}
		dfs(nd.Left)
		ans = append(ans, nd.Val)
		dfs(nd.Right)
	}
	dfs(root)
	return ans
}

func inorderTraversal2(root *TreeNode) []int { // by iter
	var part []*TreeNode
	ans, cur := make([]int, 0), root
	for cur != nil || len(part) != 0 {
		for cur != nil {
			cur, part = cur.Left, append(part, cur)
		}
		u := part[len(part)-1]
		part = part[:len(part)-1]
		ans = append(ans, u.Val)
		cur = u.Right
	}
	return ans
}

func inorderTraversal(root *TreeNode) []int { // by morris O(1) space
	ans, cur := make([]int, 0), root
	for cur != nil {
		pre := cur.Left
		if pre != nil {
			for pre.Right != cur && pre.Right != nil {
				pre = pre.Right
			}
			if pre.Right == nil { // 1st: attach
				pre.Right = cur
				cur = cur.Left // left first
				continue
			}
			if pre.Right == cur { // 2nd: detach
				pre.Right = nil
			}
		}
		// vis
		ans = append(ans, cur.Val)
		cur = cur.Right
	}
	return ans
}
