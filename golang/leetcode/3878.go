package main

func countGoodSubarrays1(nums []int) int64 {
	B, n := 30, len(nums)
	nxt := make([][30]int, n+1)
	for i := 0; i < B; i++ {
		nxt[n][i] = n
	}
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for b := 0; b < B; b++ {
			if v>>b&1 > 0 {
				nxt[i][b] = i
			} else {
				nxt[i][b] = nxt[i+1][b]
			}
		}
	}

	var pre [30]int
	for i := range pre {
		pre[i] = -1
	}
	last, ans := make(map[int]int), 0
	for i, v := range nums {
		l, r := -1, n
		for b := 0; b < B; b++ {
			if v>>b&1 > 0 {
				pre[b] = i
			} else {
				l, r = max(l, pre[b]), min(r, nxt[i+1][b])
			}
		}
		if p, ok := last[v]; ok {
			l = max(l, p)
		}
		//fmt.Println(i, last[v], v, l, r)
		ans += (i - l) * (r - i)
		last[v] = i
	}
	return int64(ans)
}

func countGoodSubarrays(nums []int) int64 { // O(n) mono-que
	n := len(nums)
	rs := make([]int, n)
	q := [][2]int{{(1 << 32) - 1, n}}
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for v|q[len(q)-1][0] == v {
			q = q[:len(q)-1]
		}
		rs[i] = q[len(q)-1][1]
		q = append(q, [2]int{v, i})
	}

	ans := 0
	q = [][2]int{{(1 << 32) - 1, -1}}
	for i, v := range nums {
		for v != q[len(q)-1][0] && v|q[len(q)-1][0] == v {
			q = q[:len(q)-1]
		}
		l := q[len(q)-1][1]
		ans += (rs[i] - i) * (i - l)
		q = append(q, [2]int{v, i})
	}
	return int64(ans)
}
