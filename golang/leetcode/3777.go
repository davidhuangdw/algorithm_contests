package main

func minDeletions(s string, queries [][]int) []int {
	n := len(s)
	fw := make([]int, n)
	add := func(i, v int) {
		for ; i < n; i += i & -i {
			fw[i] += v
		}
	}
	get := func(i int) (s int) {
		for ; i > 0; i &= i - 1 {
			s += fw[i]
		}
		return s
	}
	cur := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			cur[i] = 1
			add(i, 1)
		}
	}
	ans := make([]int, 0)
	for _, q := range queries {
		if q[0] == 1 {
			for i := max(1, q[1]); i < min(n, q[1]+2); i++ {
				add(i, (cur[i]^1)-cur[i])
				cur[i] ^= 1
			}
		} else {
			l, r := q[1], q[2]
			ans = append(ans, (r-l)-(get(r)-get(l)))
		}
	}
	return ans
}
