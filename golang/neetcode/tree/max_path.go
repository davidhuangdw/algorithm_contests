package tree

func maxPathSum(root *TreeNode) int {
	ans := int(-1e9)
	var subMax func(*TreeNode) int
	subMax = func(nd *TreeNode) int {
		if nd == nil {
			return 0
		}
		l, r := subMax(nd.Left), subMax(nd.Right)
		ans = max(ans, nd.Val+l+r)
		return max(0, nd.Val+max(l, r))
	}
	subMax(root)
	return ans
}
