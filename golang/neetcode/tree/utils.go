package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Helper function to create a binary tree from slice representation
// Uses level order traversal where nil represents missing nodes
func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			current.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			current.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}
