package main

func validSubarraySize(nums []int, th int) int {
	n := len(nums)
	st, rs := [][2]int{{0, n}}, make([]int, n)
	for i := len(nums) - 1; i >= 0; i-- {
		v := nums[i]
		for ; st[len(st)-1][0] >= v; st = st[:len(st)-1] {
		}
		rs[i] = st[len(st)-1][1]
		st = append(st, [2]int{v, i})
	}
	st = [][2]int{{0, -1}}
	for i, v := range nums {
		for ; st[len(st)-1][0] >= v; st = st[:len(st)-1] {
		}
		k := rs[i] - 1 - st[len(st)-1][1]
		if v > th/k {
			return k
		}
		st = append(st, [2]int{v, i})
	}
	return -1
}
