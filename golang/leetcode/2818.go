package main

import "slices"

func maximumScore(nums []int, k int) int {
	n, mx := len(nums), slices.Max(nums)
	score := make([]int, mx+1)
	for v := 2; v <= mx; v++ {
		if score[v] == 0 {
			score[v] = 1
			for x := v + v; x <= mx; x += v {
				score[x]++
			}
		}
	}

	ls, rs, inf := make([]int, n), make([]int, n), int(1e9)
	q := [][2]int{{inf, -1}}
	for i := 0; i < n; i++ {
		v := score[nums[i]]
		for ; q[len(q)-1][0] < v; q = q[:len(q)-1] {
		}
		ls[i] = q[len(q)-1][1]
		q = append(q, [2]int{v, i})
	}
	q = [][2]int{{inf, n}}
	for i := n - 1; i >= 0; i-- {
		v := score[nums[i]]
		for ; q[len(q)-1][0] <= v; q = q[:len(q)-1] {
		}
		rs[i] = q[len(q)-1][1]
		q = append(q, [2]int{v, i})
	}

	ord := make([]int, n)
	for i := range ord {
		ord[i] = i
	}
	slices.SortFunc(ord, func(a, b int) int {
		return nums[b] - nums[a]
	})

	M := int(1e9 + 7)
	pow := func(x, p int) int {
		res := 1
		for b := x; p > 0; p >>= 1 {
			if p&1 > 0 {
				res = res * b % M
			}
			b = b * b % M
		}
		return res
	}

	ans := 1
	for _, i := range ord {
		x, p := nums[i], min(k, (rs[i]-i)*(i-ls[i]))
		ans = ans * pow(x, p) % M
		k -= p
		if k == 0 {
			break
		}
	}
	return ans
}
