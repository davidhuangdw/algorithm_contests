package main

import (
	"slices"
	"sort"
)

func maxNumOfSubstrings1(s string) []string {
	var pos [26][]int
	for i, c := range []byte(s) {
		t := int(c - 'a')
		pos[t] = append(pos[t], i)
	}
	var st, ed [26]int
	for i := range pos {
		st[i], ed[i] = -1, -1
		if len(pos[i]) > 0 {
			st[i], ed[i] = pos[i][0], pos[i][len(pos[i])-1]
		}
	}

	var contain [26][]int
	for x := range 26 {
		if len(pos[x]) == 0 {
			continue
		}
		for y := range 26 {
			if x == y || len(pos[y]) == 0 {
				continue
			}
			j := sort.Search(len(pos[y]), func(i int) bool {
				return pos[y][i] >= st[x]
			})
			if j < len(pos[y]) && pos[y][j] <= ed[x] {
				contain[x] = append(contain[x], y)
			}
		}
	}
	var segs [][2]int
	var vis [26]bool
	var dfs func(u int) (l, r int)
	dfs = func(u int) (l, r int) {
		vis[u] = true
		l, r = st[u], ed[u]
		for _, v := range contain[u] {
			if !vis[v] {
				vl, vr := dfs(v)
				l, r = min(l, vl), max(r, vr)
			}
		}
		return
	}
	for x := range 26 {
		if st[x] >= 0 {
			vis = [26]bool{}
			l, r := dfs(x)
			segs = append(segs, [2]int{l, r})
		}

	}

	slices.SortFunc(segs, func(a, b [2]int) int {
		if a[1] != b[1] {
			return a[1] - b[1]
		}
		return b[0] - a[0]
	})
	var ans []string
	last := -1
	for _, seg := range segs {
		if seg[0] > last {
			ans = append(ans, s[seg[0]:seg[1]+1])
			last = seg[1]
		}
	}
	return ans
}

func maxNumOfSubstrings(s string) []string {
	bs := []byte(s)
	var st, ed [26]int
	for i := range st {
		st[i] = -1
	}
	for i, c := range bs {
		t := int(c - 'a')
		if st[t] < 0 {
			st[t] = i
		}
		ed[t] = i
	}

	var segs [][2]int
outer:
	for x := range 26 {
		if st[x] < 0 {
			continue
		}
		l, r := st[x], ed[x]
		for i := l; i <= r; i++ {
			t := int(bs[i] - 'a')
			if st[t] < l {
				continue outer
			}
			r = max(r, ed[t])
		}
		segs = append(segs, [2]int{l, r})
	}

	slices.SortFunc(segs, func(a, b [2]int) int {
		if a[1] != b[1] {
			return a[1] - b[1]
		}
		return b[0] - a[0]
	})
	var ans []string
	last := -1
	for _, seg := range segs {
		if seg[0] > last {
			ans = append(ans, s[seg[0]:seg[1]+1])
			last = seg[1]
		}
	}
	return ans
}
