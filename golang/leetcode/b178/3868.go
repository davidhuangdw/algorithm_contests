package main

func minCost(nums1 []int, nums2 []int) int {
	cnt := make(map[int]int)
	for _, v := range nums1 {
		cnt[v]++
	}
	for _, v := range nums2 {
		cnt[v]--
	}
	ans := 0
	for _, k := range cnt {
		if k < 0 {
			k = -k
		}
		if k%2 == 1 {
			return -1
		}
		ans += k / 2
	}
	return ans / 2
}
