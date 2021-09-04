package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	var ret [][]int
	sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	l := intervals[0][0]
	r := intervals[0][1]
	for _, rng := range intervals{
		if rng[0] <= r {
			if r < rng[1]{
				r = rng[1]
			}
		}else{
			ret = append(ret, []int{l, r})
			l = rng[0]
			r = rng[1]
		}
	}
	ret = append(ret, []int{l, r})
	return ret
}