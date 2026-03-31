package main

import "math/bits"

func maxStudents(seats [][]byte) int {
	n, m := len(seats), len(seats[0])
	row, M := make([]int, n), 1<<m
	for i, s := range seats {
		for j, c := range s {
			if c == '.' {
				row[i] |= 1 << j
			}
		}
	}

	dp, pre, msk := make([]int, M), 0, (1<<M)-1
	for _, seat := range row {
		ndp := make([]int, M)
		ndp[0] = dp[pre]
		for allow := seat; allow > 0; allow = (allow - 1) & seat {
			ndp[allow] = dp[pre] // cur==0
			for cur := allow; cur > 0; cur = (cur - 1) & allow {
				forbid := (cur << 1) | (cur >> 1)
				k := bits.OnesCount(uint(cur))
				if (forbid)&cur == 0 {
					ndp[allow] = max(ndp[allow], k+dp[pre&(msk^forbid)])
				}
			}
		}
		dp, pre = ndp, seat
	}
	return dp[row[n-1]]
}
