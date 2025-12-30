package intervals

func insert(origin [][]int, newInterval []int) [][]int {
	l, r, n := newInterval[0], newInterval[1], len(origin)
	i, ans := 0, make([][]int, 0)
	for ; i < n && origin[i][1] < l; i++ {
		ans = append(ans, origin[i])
	}
	for ; i < n && origin[i][0] <= r; i++ {
		l, r = min(l, origin[i][0]), max(r, origin[i][1])
	}
	ans = append(ans, []int{l, r})
	return append(ans, origin[i:]...)
}
