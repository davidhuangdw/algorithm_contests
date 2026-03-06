package main

import "math/rand/v2"

type TreapNode struct {
	v, pri  int
	l, r    *TreapNode
	rep, sz int
}

func NewTreapNode(x, p int) *TreapNode {
	if p <= 0 {
		p = rand.Int()
	}
	return &TreapNode{v: x, pri: p, l: nil, r: nil, sz: 1, rep: 1}
}
func (nd *TreapNode) Update() {
	if nd == nil {
		return
	}
	nd.sz = nd.rep
	for _, ch := range []*TreapNode{nd.l, nd.r} {
		if ch != nil {
			nd.sz += ch.sz
		}
	}
}
func (nd *TreapNode) Size() int {
	if nd == nil {
		return 0
	}
	return nd.sz
}

type FhqTreap struct {
	root *TreapNode
}

func (tr *FhqTreap) Split(u *TreapNode, x int) (L, R *TreapNode) { // <=, >
	defer u.Update()
	if u == nil {
		return nil, nil
	}
	if u.v <= x {
		u.r, R = tr.Split(u.r, x)
		return u, R
	} else {
		L, u.l = tr.Split(u.l, x)
		return L, u
	}
}

func (tr *FhqTreap) Merge(L, R *TreapNode) *TreapNode {
	if L == nil {
		return R
	}
	if R == nil {
		return L
	}
	if L.pri <= R.pri {
		L.r = tr.Merge(L.r, R)
		L.Update()
		return L
	} else {
		R.l = tr.Merge(L, R.l)
		R.Update()
		return R
	}
}

func (tr *FhqTreap) Insert(x int) {
	L, R := tr.Split(tr.root, x)
	L, eq := tr.Split(L, x-1)
	if eq != nil {
		eq.rep += 1
		eq.sz += 1
	} else {
		eq = NewTreapNode(x, -1)
	}
	tr.root = tr.Merge(tr.Merge(L, eq), R)
}

func (tr *FhqTreap) Del(x int) {
	L, R := tr.Split(tr.root, x)
	L, eq := tr.Split(L, x-1)
	if eq != nil {
		eq.rep--
		eq.sz--
		if eq.rep == 0 {
			eq = tr.Merge(eq.l, eq.r)
		}
	}
	tr.root = tr.Merge(tr.Merge(L, eq), R)
}
func (tr *FhqTreap) Rank(x int) int { // count(<x)+1
	L, R := tr.Split(tr.root, x-1)
	ans := 1
	if L != nil {
		ans += L.sz
	}
	tr.root = tr.Merge(L, R)
	return ans
}
func (tr *FhqTreap) Kth(k int) int {
	return tr.SubKth(tr.root, k)
}

func (tr *FhqTreap) SubKth(u *TreapNode, k int) int {
	for u != nil && k > 0 {
		if k <= u.l.Size() {
			return tr.SubKth(u.l, k)
		} else if k <= u.l.Size()+u.rep {
			return u.v
		} else {
			k -= u.l.Size() + u.rep
			u = u.r
		}
	}
	return -1
}

func (tr *FhqTreap) Prev(x int) int {
	L, R := tr.Split(tr.root, x-1)
	defer func() { tr.root = tr.Merge(L, R) }()
	return tr.SubKth(L, L.Size())
}
func (tr *FhqTreap) Succ(x int) int {
	L, R := tr.Split(tr.root, x)
	defer func() { tr.root = tr.Merge(L, R) }()
	return tr.SubKth(R, 1)
}
