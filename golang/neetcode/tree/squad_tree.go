package tree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var dfs func(i, j, n int) *Node
	dfs = func(i, j, n int) *Node {
		if n == 0 {
			return nil
		}
		if n == 1 {
			return &Node{Val: grid[i][j] == 1, IsLeaf: true}
		}
		n /= 2
		ch, same := []*Node{dfs(i, j, n), dfs(i, j+n, n), dfs(i+n, j, n), dfs(i+n, j+n, n)}, true
		for t := 0; same && t < 3; t++ {
			same = same && (ch[t].IsLeaf && ch[t+1].IsLeaf && ch[t].Val == ch[t+1].Val)
		}
		if same {
			return ch[0]
		}
		return &Node{
			IsLeaf:      false,
			TopLeft:     ch[0],
			TopRight:    ch[1],
			BottomLeft:  ch[2],
			BottomRight: ch[3],
		}
	}
	return dfs(0, 0, len(grid))
}
