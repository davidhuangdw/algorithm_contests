package graphs

import (
	"slices"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	eiOrd := make([]int, len(edges))
	for i := range eiOrd {
		eiOrd[i] = i
	}
	slices.SortFunc(eiOrd, func(a, b int) int {
		return edges[a][2] - edges[b][2]
	})
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var root func(i int) int
	root = func(i int) int {
		if fa[i] != i {
			fa[i] = root(fa[i])
		}
		return fa[i]
	}

	mstEdges := make(map[int]bool) // mstEdges-edges
	notCri := make(map[int]bool)
	mstNxt := make([][][]int, n)
	ans := make([][]int, 2)
	for _, i := range eiOrd {
		e := edges[i]
		u, v := e[0], e[1]
		ru, rv := root(u), root(v)
		if ru == rv {
			continue
		}
		mstEdges[i] = true
		fa[ru] = rv

		mstNxt[u] = append(mstNxt[u], []int{v, i})
		mstNxt[v] = append(mstNxt[v], []int{u, i})
	}

	for i := 0; i < len(edges); i++ {
		if mstEdges[i] {
			continue
		}
		e := edges[i]
		st, ed, target := e[0], e[1], e[2]
		var findPath func(u, par int, path []int) bool
		findPath = func(u, par int, path []int) bool {
			if u == ed {
				for _, ei := range path {
					if edges[ei][2] == target {
						notCri[ei] = true
						notCri[i] = true
					}
				}
				return true
			}
			for _, nx := range mstNxt[u] {
				v, ei := nx[0], nx[1]
				if v == par {
					continue
				}
				if findPath(v, u, append(path, ei)) {
					return true
				}
			}
			return false
		}
		findPath(st, -1, nil)
	}

	for ei := range mstEdges {
		if !notCri[ei] {
			ans[0] = append(ans[0], ei)
		}
	}
	for ei := range notCri {
		ans[1] = append(ans[1], ei)
	}
	return ans
}
