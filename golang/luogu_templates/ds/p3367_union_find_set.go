package main

import (
	"fmt"
	"io"
)

type UnionFindSet struct {
	parent []int
	rank   []int
}

func NewUnionFindSet(n int) *UnionFindSet {
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := 0; i <= n; i++ {
		rank[i] = 1
		parent[i] = i
	}
	return &UnionFindSet{
		parent: parent,
		rank:   rank,
	}
}

func (uf *UnionFindSet) find(u int) int {
	if uf.parent[u] != u {
		uf.parent[u] = uf.find(uf.parent[u])
	}
	return uf.parent[u]
}

func (uf *UnionFindSet) union(u, v int) {
	if u == v {
		return
	}
	a, b := uf.find(u), uf.find(v)
	if a == b {
		return
	}
	if uf.rank[a] > uf.rank[b] {
		a, b = b, a
	}
	uf.parent[b] = a
	uf.rank[a] += uf.rank[b]
}

func P3367_union_find_set(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	s := NewUnionFindSet(n)
	for i := 0; i < m; i++ {
		var z, x, y int
		fmt.Fscan(IN, &z, &x, &y)
		switch z {
		case 1:
			s.union(x, y)
		case 2:
			if s.find(x) == s.find(y) {
				fmt.Fprintln(OUT, "Y")
			} else {
				fmt.Fprintln(OUT, "N")
			}
		}
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3367_union_find_set(IN, OUT)
//}
