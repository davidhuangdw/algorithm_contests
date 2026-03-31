package main

import (
	"math/bits"
	"slices"
)

func minimumIncompatibility1(nums []int, k int) int {
	n := len(nums)
	slices.Sort(nums)
	mx, sz, inf := 1<<n, n/k, int(1e15)
	sub := make([]int, mx)
	for i := range sub {
		sub[i] = -1
	}
outer:
	for msk := 1; msk < mx; msk++ {
		if bits.OnesCount(uint(msk)) == sz {
			l, r, pre := inf, 0, -1
			for i, v := range nums {
				if msk&(1<<i) > 0 {
					if v == pre {
						continue outer
					}
					l, r, pre = min(l, v), max(r, v), v
				}
			}
			sub[msk] = r - l
		}
	}

	dp := make([]int, mx)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for msk := 0; msk < mx-1; msk++ {
		if dp[msk] == inf || bits.OnesCount(uint(msk)) > n-sz {
			continue
		}
		oth := (mx - 1) ^ msk
		for nxt := oth; nxt > 0; nxt = (nxt - 1) & oth {
			if sub[nxt] >= 0 {
				dp[msk|nxt] = min(dp[msk|nxt], dp[msk]+sub[nxt])
				//fmt.Printf("%b: %b: %v\n", msk, nxt, dp[msk|nxt])
			}
		}
	}
	if dp[mx-1] == inf {
		return -1
	}
	return dp[mx-1]
}

func minimumIncompatibility(nums []int, k int) int {
	inf, n, sz := int(1e12), len(nums), len(nums)/k
	slices.Sort(nums)
	mem := make(map[[2]int]int)
	var dp func(msk, last int) int
	dp = func(msk, last int) int {
		if msk == 0 {
			return 0
		}
		st := [2]int{msk, last}
		if _, ok := mem[st]; !ok {
			cnt := bits.OnesCount(uint(msk))
			if cnt%sz == 0 {
				for i := 0; i < n; i++ {
					if msk&(1<<i) > 0 {
						return dp(msk^(1<<i), i)
					}
				}
			}
			res := inf
			for i := last + 1; i < n; i++ {
				if (msk&(1<<i) > 0) && nums[i] != nums[last] {
					res = min(res, dp(msk^(1<<i), i)+nums[i]-nums[last])
				}
			}
			mem[st] = res
		}
		return mem[st]
	}
	ans := dp(1<<n-1, -1)
	if ans >= inf {
		return -1
	}
	return ans
}
