package dp

type Node struct {
	nxt map[byte]*Node
	end bool
}

func NewNode() *Node { return &Node{make(map[byte]*Node), false} }
func AddWord(nd *Node, s string) {
	for _, b := range []byte(s) {
		if _, ok := nd.nxt[b]; !ok {
			nd.nxt[b] = NewNode()
		}
		nd = nd.nxt[b]
	}
	nd.end = true
}

func wordBreak(s string, wordDict []string) bool {
	root := NewNode()
	for _, wd := range wordDict {
		AddWord(root, wd)
	}

	n := len(s)
	valid := make([]bool, n+1)
	valid[n] = true
	for i := len(s) - 1; i >= 0; i-- {
		nd, succ := root, false

		for j := i; !succ && j < n && nd != nil; j++ {
			nd = nd.nxt[s[j]]
			succ = nd != nil && nd.end && valid[j+1]
		}
		valid[i] = succ
	}
	return valid[0]
}
