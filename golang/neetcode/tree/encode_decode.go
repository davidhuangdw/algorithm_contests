package tree

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var ss []string
	var dfs func(*TreeNode)
	dfs = func(nd *TreeNode) {
		if nd == nil {
			ss = append(ss, "@")
			return
		}
		ss = append(ss, strconv.Itoa(nd.Val))
		dfs(nd.Left)
		dfs(nd.Right)
	}
	dfs(root)
	return strings.Join(ss, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	ss, i := strings.Split(data, ","), 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if ss[i] == "@" {
			i++
			return nil
		}
		v, _ := strconv.Atoi(ss[i])
		i++
		return &TreeNode{v, build(), build()}
	}
	return build()
}
