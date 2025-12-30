package intervals

import "slices"

func merge1(intervals [][]int) [][]int {
	ans, n := make([][]int, 0), len(intervals)
	if n == 0 {
		return ans
	}
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	l, r := intervals[0][0], intervals[0][1]
	for _, p := range intervals[1:] {
		if r < p[0] {
			ans, l, r = append(ans, []int{l, r}), p[0], p[1]
		} else {
			r = max(r, p[1])
		}
	}
	return append(ans, []int{l, r})
}

func merge(intervals [][]int) [][]int { // O(val) by iter each val
	ans, n := make([][]int, 0), len(intervals)
	if n == 0 {
		return ans
	}
	byLeft, mn, mx := make(map[int]int), int(1e9), int(-1e9)
	for _, p := range intervals {
		if r, ok := byLeft[p[0]]; !ok || r < p[1] { // bug: might be negative
			byLeft[p[0]] = p[1]
		}
		mx = max(mx, p[0])
		mn = min(mn, p[0])
	}

	l, r := mn, byLeft[mn]
	for i := mn + 1; i <= mx; i++ {
		j, ok := byLeft[i]
		if !ok {
			continue
		}
		if r < i {
			ans, l, r = append(ans, []int{l, r}), i, j
		} else {
			r = max(r, j)
		}
	}
	return append(ans, []int{l, r})
}
