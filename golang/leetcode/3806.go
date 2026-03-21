package main

import "slices"

func maximumAND(nums []int, k int, m int) int {
	want := 0
	dis := make([]int, len(nums))
	for b := 30; b >= 0; b-- {
		want += 1 << b
		for i, v := range nums {
			dis[i] = 0
			if v&want != want {
				msk := want & ^v
				for msk&(msk-1) > 0 {
					msk &= msk - 1
				}
				msk = (msk << 1) - 1
				dis[i] = (want & msk) - (v & msk)
			}
		}
		s := 0
		slices.Sort(dis)
		for _, d := range dis[:m] {
			s += d
		}
		if s > k {
			want -= 1 << b
		}
	}
	return want
}
