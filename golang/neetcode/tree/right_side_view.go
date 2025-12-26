package tree

func rightSideView1(root *TreeNode) (ans []int) { // bfs
	if root == nil {
		return
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		ans = append(ans, que[0].Val)
		var nxt []*TreeNode
		for _, nd := range que {
			if nd.Right != nil {
				nxt = append(nxt, nd.Right)
			}
			if nd.Left != nil {
				nxt = append(nxt, nd.Left)
			}
		}
		que = nxt
	}
	return
}

func rightSideView(root *TreeNode) (ans []int) { // dfs, depth as index
	var dfs func(*TreeNode, int)
	dfs = func(nd *TreeNode, d int) {
		if nd == nil {
			return
		}
		if len(ans) == d {
			ans = append(ans, nd.Val)
		}
		dfs(nd.Right, d+1)
		dfs(nd.Left, d+1)
	}
	dfs(root, 0)
	return
}
