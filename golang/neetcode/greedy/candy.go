package greedy

// idea: each valley is 1
func candy(ratings []int) int {
	ans, n := 0, len(ratings)
	for i := 0; i < n; {
		// i is valley: find next valley, and record top
		l, top := i, i
		for top+1 < n && ratings[top] < ratings[top+1] {
			top++
		}
		r := top
		for r+1 < n && ratings[r] > ratings[r+1] {
			r++
		}
		ltop, rtop := top-l+1, r-top+1

		ans += (1+ltop)*ltop/2 + (1+rtop)*rtop/2
		ans -= min(ltop, rtop)

		if i == r {
			i++
		} else {
			ans-- // revert valley r for next round
			i = r
		}
	}
	return ans
}
