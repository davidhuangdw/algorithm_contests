package tree

func kthSmallest1(root *TreeNode, k int) (ans int) { // by recur
	var inorder func(*TreeNode) bool
	inorder = func(nd *TreeNode) bool {
		if nd == nil {
			return false
		}
		if ok := inorder(nd.Left); ok {
			return true
		}
		k--
		if k == 0 {
			ans = nd.Val
			return true
		}
		return inorder(nd.Right)
	}
	inorder(root)
	return
}

func kthSmallest2(root *TreeNode, k int) (ans int) { // by iter
	var hf []*TreeNode
	for len(hf) > 0 || root != nil {
		for root != nil {
			hf = append(hf, root)
			root = root.Left
		}
		root, hf = hf[len(hf)-1], hf[:len(hf)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}

func kthSmallest(root *TreeNode, k int) (ans int) { // O(1) space by morris-traversal
	for root != nil {
		if root.Left != nil {
			// check prev
			prev := root.Left
			for prev.Right != nil && prev.Right != root {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = root
				root = root.Left
				continue
			} else {
				prev.Right = nil
			}
		}
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}
