package main

import (
	"slices"
	"sort"
)

func findMinimumTime1(tasks [][]int) int {
	slices.SortFunc(tasks, func(a, b []int) int {
		return a[1] - b[1]
	})
	vis, ans := make([]bool, 2001), 0
	for _, t := range tasks {
		l, r, d := t[0], t[1], t[2]
		for i := l; i <= r; i++ {
			if vis[i] {
				d--
			}
		}
		for i := r; d > 0; i-- {
			if !vis[i] {
				vis[i] = true
				d--
				ans++
			}
		}
	}
	return ans
}

func findMinimumTime(tasks [][]int) int { // O(n*log) binary-search mono-stack
	slices.SortFunc(tasks, func(a, b []int) int {
		return a[1] - b[1]
	})
	st := [][3]int{{-2, -3, 0}} // [3]int := {l,r,preSum}
	for _, t := range tasks {
		l, r, d := t[0], t[1], t[2]
		fst := sort.Search(len(st), func(i int) bool {
			return l <= st[i][1] // first interval that overlaps or follows l
		})
		if fst < len(st) {
			d -= st[len(st)-1][2] - st[fst][2] + (st[fst][1] + 1 - max(st[fst][0], l)) // prev-covered
		}
		if d > 0 { // bug: maybe <=0
			l = r + 1 - d
			for st[len(st)-1][1]+1 >= l {
				l -= st[len(st)-1][1] + 1 - st[len(st)-1][0]
				st = st[:len(st)-1]
			}
			st = append(st, [3]int{l, r, st[len(st)-1][2] + r - l + 1})
		}
	}
	return st[len(st)-1][2]
}
