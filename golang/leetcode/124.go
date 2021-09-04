package leetcode

func maxPathSum(root *TreeNode) int {
	mx := root.Val
	max := func(a int, b int) int {
		if a < b {
			return b
		}
		return a
	}
	var dfs func(u *TreeNode) int
	dfs = func(u *TreeNode) int {
		if u == nil {
			return 0
		}
		l := max(0, dfs(u.Left))
		r := max(0, dfs(u.Right))
		mx = max(mx, u.Val+l+r)
		return u.Val + max(l, r)
	}

	dfs(root)
	return mx
}
