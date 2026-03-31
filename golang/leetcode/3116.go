package main

import "math/bits"

func findKthSmallest(coins []int, k int) int64 {
	n := len(coins)
	M := 1 << n
	lcm := make([]int, M)
	lcm[0] = 1
	gcd := func(a, b int) int {
		for a > 0 {
			a, b = b%a, a
		}
		return b
	}
	for i, v := range coins {
		for j := 0; j < (1 << i); j++ {
			lcm[j|(1<<i)] = v / gcd(lcm[j], v) * lcm[j]
		}
	}

	l, r := 1, coins[0]*k
	for l <= r {
		md, cnt := (l+r)>>1, 0
		for i := 1; i < M; i++ {
			if bits.OnesCount(uint(i))%2 == 0 {
				cnt -= md / lcm[i]
			} else {
				cnt += md / lcm[i]
			}
		}
		if cnt >= k {
			r = md - 1
		} else {
			l = md + 1
		}
	}
	return int64(l)
}
