package tree

func rob(root *TreeNode) int {
	var dfs func(nd *TreeNode) (int, int)
	dfs = func(nd *TreeNode) (a int, b int) {
		if nd == nil {
			return
		}
		la, lb := dfs(nd.Left)
		ra, rb := dfs(nd.Right)
		b = la + ra
		a = max(b, nd.Val+lb+rb)
		return a, b
	}
	return max(dfs(root))
}
