package dp

func lengthOfLIS(nums []int) int {
	INF, n := int(1e9), len(nums)
	lenMin, mx := make([]int, n+1), 0
	for i := range lenMin {
		lenMin[i] = INF
	}
	lenMin[0] = -INF

	for _, v := range nums {
		l, r := 0, mx
		for l <= r {
			md := (l + r) / 2
			if lenMin[md] < v {
				l = md + 1
			} else {
				r = md - 1
			}
		}
		lenMin[l] = min(lenMin[l], v)
		mx = max(mx, l)
	}
	return mx
}
