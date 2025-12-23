package stack

func dailyTemperatures1(t []int) []int { // by candidates-stack
	n := len(t)
	ans := make([]int, n)
	var cand []int // higher candidates
	for i := n - 1; i >= 0; i-- {
		v := t[i]
		for len(cand) > 0 && v >= t[cand[len(cand)-1]] {
			cand = cand[:len(cand)-1] // pop
		}
		if len(cand) > 0 {
			ans[i] = cand[len(cand)-1] - i
		}
		cand = append(cand, i)
	}
	return ans
}

func dailyTemperatures(t []int) []int { // O(1) extra space by follow higher-link
	n := len(t)
	higher := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		v, j := t[i], i+1
		for j < n && v >= t[j] { // total time O(n), since each j is skipped only once
			j = higher[j]
		}
		higher[i] = j
	}
	for i, j := range higher { // convert to diff
		if j == n {
			higher[i] = 0
		} else {
			higher[i] -= i
		}
	}
	return higher
}
