package main

import "strings"

func removeSubfolders(folder []string) []string {
	type Node struct {
		nxt map[string]*Node
		ed  bool
	}
	newNode := func() *Node {
		return &Node{nxt: make(map[string]*Node)}
	}
	root := newNode()
	add := func(path string) {
		p := root
		for _, s := range strings.Split(path[1:], "/") {
			if p.nxt[s] == nil {
				p.nxt[s] = newNode()
			}
			p = p.nxt[s]
			if p.ed {
				return
			}
		}
		p.ed = true
	}

	for _, f := range folder {
		add(f)
	}
	ans := make([]string, 0)
	var dfs func(p *Node, pre []string)
	dfs = func(p *Node, pre []string) {
		if p.ed {
			ans = append(ans, "/"+strings.Join(pre, "/"))
			return
		}
		for s, q := range p.nxt {
			dfs(q, append(pre, s))
		}
	}
	dfs(root, nil)
	return ans
}
