package tree

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	a, b := min(p.Val, q.Val), max(p.Val, q.Val)
	for root != nil && !(a <= root.Val && root.Val <= b) {
		if b < root.Val {
			root = root.Left
		} else if root.Val < a {
			root = root.Right
		}
	}
	return root
}
