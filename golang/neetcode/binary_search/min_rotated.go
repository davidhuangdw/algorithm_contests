package binary_search

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1 // >=, <
	for l <= r {
		md := (l + r) / 2
		if nums[md] >= nums[0] {
			l = md + 1
		} else {
			r = md - 1
		}
	}
	return nums[l%n]
}
