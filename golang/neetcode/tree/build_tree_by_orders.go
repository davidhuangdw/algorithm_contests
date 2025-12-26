package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	i, j, m := 0, 0, len(inorder)
	INF := int(1e9)
	var buildTill func(int) *TreeNode
	buildTill = func(parent int) *TreeNode {
		if j >= m {
			return nil
		}
		if inorder[j] == parent {
			j++
			return nil
		}
		root := &TreeNode{Val: preorder[i]}
		i++
		root.Left = buildTill(root.Val)
		root.Right = buildTill(parent)
		return root
	}
	return buildTill(INF)
}