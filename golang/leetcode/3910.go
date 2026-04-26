package main

import "math/bits"

func evenSumSubgraphs(nums []int, edges [][]int) (ans int) {
	n := len(nums)
	adj := make([]int, n)
	for _, e := range edges {
		adj[e[0]] |= 1 << e[1]
		adj[e[1]] |= 1 << e[0]
	}

	all := 0
	for i, v := range nums {
		all |= v << i
	}

	for msk := 1; msk < 1<<n; msk++ {
		if bits.OnesCount(uint(msk&all))%2 != 0 {
			continue
		}
		vis := msk & -msk
		for prev := 0; vis != prev; {
			prev = vis
			for i := range n {
				if vis&(1<<i) != 0 {
					vis |= adj[i] & msk
				}
			}
		}
		if vis == msk {
			ans++
		}
	}
	return ans
}
