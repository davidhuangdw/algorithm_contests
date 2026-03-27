package main

func closestToTarget(arr []int, target int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	ans := int(1e9)
	var q []int
	for _, v := range arr {
		q = append(q, v)
		j := 0
		for _, x := range q {
			x &= v
			if j == 0 || q[j-1] != x {
				q[j] = x
				j++
				ans = min(ans, abs(target-x))
			}
		}
		q = q[:j]
	}
	return ans
}
