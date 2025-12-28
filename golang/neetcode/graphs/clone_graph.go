package graphs

func cloneGraph1(node *Node) *Node {
	if node == nil {
		return nil
	}
	clones := make(map[*Node]*Node)
	var dfs func(parent *Node) *Node
	dfs = func(parent *Node) *Node {
		if clones[parent] != nil {
			return clones[parent]
		}
		cl := &Node{Val: parent.Val}
		clones[parent] = cl
		for _, ch := range parent.Neighbors {
			cl.Neighbors = append(cl.Neighbors, dfs(ch))
		}
		return cl
	}
	return dfs(node)
}

func cloneGraph(node *Node) *Node { // bfs
	if node == nil {
		return nil
	}
	clones := make(map[*Node]*Node)
	que := []*Node{node}
	clones[node] = &Node{Val: node.Val} // qued before <-> inside clones[]
	for len(que) > 0 {
		nd := que[0]
		que = que[1:]
		cl := clones[nd]
		for _, ne := range nd.Neighbors {
			if _, ok := clones[ne]; !ok {
				clones[ne] = &Node{Val: ne.Val}
				que = append(que, ne)
			}
			cl.Neighbors = append(cl.Neighbors, clones[ne])
		}
	}
	return clones[node]
}
