package tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	var merge func(l, r *TreeNode) *TreeNode
	merge = func(l, r *TreeNode) *TreeNode {
		if l == nil {
			return r
		}
		if r == nil {
			return l
		}
		r.Left = merge(l, r.Left)
		return r
	}

	if key == root.Val {
		return merge(root.Left, root.Right)
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
