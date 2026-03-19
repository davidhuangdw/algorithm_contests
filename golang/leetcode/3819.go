package main

func rotateElements(nums []int, k int) []int {
	var ps []int
	for i, v := range nums {
		if v >= 0 {
			ps = append(ps, i)
		}
	}
	n := len(ps)
	if n == 0 {
		return nums
	}
	d, c := n-k%n, 0
	for st := 0; st < n && c < n; st++ {
		for j := (st + d) % n; j != st; j = (j + d) % n {
			nums[ps[st]], nums[ps[j]] = nums[ps[j]], nums[ps[st]]
			c++
		}
		c++
	}
	return nums
}
