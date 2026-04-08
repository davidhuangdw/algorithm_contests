package main

func minIncrease(nums []int) int64 {
	n := len(nums)
	cost := func(i int) int {
		return max(0, max(nums[i-1], nums[i+1])+1-nums[i])
	}

	suf := 0
	for i := n - 2; i > 0; i -= 2 {
		suf += cost(i)
	}
	if n%2 == 1 {
		return int64(suf)
	}

	ans, cur := suf, suf
	for i := 1; i+2 < n; i += 2 {
		cur += cost(i) - cost(i+1)
		ans = min(ans, cur)
	}
	return int64(ans)
}
