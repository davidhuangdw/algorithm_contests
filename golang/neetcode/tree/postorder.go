package tree

import "slices"

func postorderTraversal1(root *TreeNode) []int { // by iter
	var part [][]*TreeNode
	cur, ans := root, make([]int, 0)
	for !(cur == nil && len(part) == 0) {
		for cur != nil {
			part = append(part, []*TreeNode{cur.Right, cur})
			cur = cur.Left
		}
		p := part[len(part)-1]
		part = part[:len(part)-1]
		if p[0] == nil { // vis
			ans = append(ans, p[1].Val)
		} else { // right not done
			cur = p[0]
			part = append(part, []*TreeNode{nil, p[1]})
		}
	}
	return ans
}

func postorderTraversal(root *TreeNode) []int { // by morris & reverse
	cur, ans := root, make([]int, 0)
	for cur != nil {
		if cur.Right == nil { // one-side
			ans, cur = append(ans, cur.Val), cur.Left
			continue
		}
		pre := cur.Right
		for pre.Left != nil && pre.Left != cur {
			pre = pre.Left
		}
		if pre.Left == nil { // right undone
			pre.Left = cur // attach
			ans = append(ans, cur.Val)
			cur = cur.Right
		} else { // right done
			pre.Left = nil
			cur = cur.Left
		}
	}

	slices.Reverse(ans)
	return ans
}
