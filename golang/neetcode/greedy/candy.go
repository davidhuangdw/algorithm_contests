package greedy

// idea: each valley is 1
func candy1(ratings []int) int {
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

		ans += (1+ltop)*ltop/2 + (1+rtop)*rtop/2 - min(ltop, rtop)

		if i == r {
			i++
		} else {
			ans-- // revert valley r for next round
			i = r
		}
	}
	return ans
}

func candy(ratings []int) int {
	// optimize cognitive: 1) valley as 0 by ans begin with n, 2) skip empty-mountain(r[i]==ra[i+1])
	n := len(ratings)
	ans := n             // valley as 0
	for i := 0; i < n; { //  i is valley, mountain := i~next-valley
		if i+1 == n || ratings[i+1] == ratings[i] { // empty-mountain
			i++
			continue
		}
		// non-empty mountain:
		up, down := 0, 0
		j := i + 1
		for ; j < n && ratings[j-1] < ratings[j]; j++ {
			up++
		}
		for ; j < n && ratings[j-1] > ratings[j]; j++ {
			down++
		}
		ans += (1+up)*up/2 + (1+down)*down/2 - min(up, down)
		i = j - 1 // next valley
	}
	return ans
}
