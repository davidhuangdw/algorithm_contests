package two_pointers

import "slices"

func fourSum(nums []int, target int) [][]int {
	n, ans := len(nums), make([][]int, 0)
	slices.Sort(nums)
	for a, va := range nums {
		if a-1 >= 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < n-2; b++ {
			if b-1 > a && nums[b] == nums[b-1] {
				continue
			}
			d, vb := n-1, nums[b]
			for c := b + 1; c < d; c++ {
				if c-1 > b && nums[c] == nums[c-1] {
					continue
				}
				vc := nums[c]
				vd := target - va - vb - vc
				for c < d && nums[d] >= vd {
					d--
				}
				if d+1 < n && nums[d+1] == vd {
					ans = append(ans, []int{va, vb, vc, vd})
				}
			}
		}
	}
	return ans
}
