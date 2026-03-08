package main

func smallestBalancedIndex(nums []int) int {
	n := len(nums)
	rs := make([]int, n+1) // nums[i]>0
	rs[n] = 1
	for i := n - 1; i >= 0; i-- {
		mul := nums[i] * rs[i+1]
		if nums[i] == 0 || mul/nums[i] == rs[i+1] {
			rs[i] = mul
		} else { // overflow
			break
		}
	}
	pre := 0
	for i, v := range nums {
		if i > 0 && pre == rs[i+1] {
			return i
		}
		pre += v
	}
	return -1
}
