package intervals

import "slices"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[1] - b[1]
	})
	cnt, r := 0, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] < r {
			cnt++
		} else {
			r = p[1]
		}
	}
	return cnt
}
