package main

func minMergeCost(lists [][]int) int64 {
	merge := func(a, b []int) []int {
		res := make([]int, 0, len(a)+len(b))
		for i, j := 0, 0; i < len(a) || j < len(b); {
			if i < len(a) && (j >= len(b) || a[i] < b[j]) {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
		return res
	}
	n := len(lists)
	arr := make([][]int, 1<<n)
	for b := 1; b < 1<<n; b++ {
		var i int
		for i = 0; i < n && (b>>i&1 == 0); i++ {
		}
		arr[b] = merge(arr[b-(1<<i)], lists[i])
	}

	med := func(a []int) int {
		return a[(len(a)-1)/2]
	}
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	dp := make([]int, 1<<n)
	for b := 1; b < 1<<n; b++ {
		mn := 0 // bug: shouldn't be -1, need allow dp[single-item]=0
		for x := 1; x < b-x; x++ {
			if x&b == x {
				cost := dp[x] + dp[b-x] + len(arr[b]) + abs(med(arr[x])-med(arr[b-x]))
				if mn <= 0 || cost < mn {
					mn = cost
				}
			}
		}
		dp[b] = mn
	}
	return int64(dp[(1<<n)-1])
}
