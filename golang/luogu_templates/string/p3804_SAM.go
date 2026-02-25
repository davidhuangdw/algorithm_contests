package main

import (
	"fmt"
	"io"
	"slices"
)

type SamNode struct {
	l, fa, cnt int
	son        map[byte]int
}

type SuffixAutomation struct {
	nodes      []SamNode
	last, root int
}

func NewSuffixAutomation() *SuffixAutomation {
	sa := &SuffixAutomation{}
	sa.root = sa.NewNode(0)
	sa.nodes[sa.root].fa = -1 // root
	sa.last = sa.root
	return sa
}

func (su *SuffixAutomation) NewNode(l int) int {
	i := len(su.nodes)
	su.nodes = append(su.nodes, SamNode{l: l, son: make(map[byte]int)})
	return i
}
func (su *SuffixAutomation) cloneNode(old int) int {
	u := su.NewNode(0)
	a := su.nodes
	a[u].l, a[u].fa = a[old].l, a[old].fa
	for k, v := range a[old].son {
		a[u].son[k] = v
	}
	return u
}

func (su *SuffixAutomation) extend(ch byte) {
	p, cur := su.last, su.NewNode(su.nodes[su.last].l+1)
	a := su.nodes
	a[cur].cnt = 1 // non-clone
	su.last = cur
	for ; p != -1 && a[p].son[ch] == 0; p = a[p].fa {
		a[p].son[ch] = cur
	}
	// p is pre of cur's father q: p+ch == q
	if p == -1 {
		a[cur].fa = su.root
		return
	}
	q := a[p].son[ch]
	if a[q].l == a[p].l+1 {
		a[cur].fa = q
		return
	}
	// q is more than 'p + ch', need split
	nq := su.cloneNode(q)
	a = su.nodes
	a[nq].l = a[p].l + 1

	a[cur].fa = nq
	a[q].fa = nq
	// all 'p+ch' change from q to nq now:
	for ; p != -1 && a[p].son[ch] == q; p = a[p].fa {
		a[p].son[ch] = nq
	}
}

func (su *SuffixAutomation) buildCnt() {
	// order by l desc
	a := su.nodes
	ord := make([]int, len(a))
	for i := range ord {
		ord[i] = i
	}
	slices.SortFunc(ord, func(i, j int) int {
		return a[j].l - a[i].l
	})

	// add to 'fa'
	for _, u := range ord {
		fa := a[u].fa
		if fa != -1 {
			a[fa].cnt += a[u].cnt
		}
	}
}

func P3804_SAM(IN io.Reader, OUT io.Writer) {
	var s string
	fmt.Fscan(IN, &s)
	sam := NewSuffixAutomation()

	// build SAM
	for _, ch := range []byte(s) {
		sam.extend(ch)
	}
	sam.buildCnt()
	mx := 0
	for _, nd := range sam.nodes {
		if nd.cnt >= 2 {
			mx = max(mx, nd.cnt*nd.l)
		}
	}
	fmt.Fprintln(OUT, mx)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3804_SAM(IN, OUT)
//}
