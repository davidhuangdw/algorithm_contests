package tree

func height(r *TreeNode) int {
	if r == nil {
		return 0
	}
	return 1 + max(height(r.Left), height(r.Right))
}
func same(l, r *TreeNode) bool {
	if l == r {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && same(l.Left, r.Left) && same(l.Right, r.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	subH := height(subRoot)
	var search func(*TreeNode) (int, bool)
	search = func(nd *TreeNode) (int, bool) {
		if nd == nil {
			return 0, false
		}
		lh, lok := search(nd.Left)
		rh, rok := search(nd.Right)
		h := 1 + max(lh, rh)
		return h, lok || rok || (h == subH && same(nd, subRoot))
	}
	_, ok := search(root)
	return ok
}
