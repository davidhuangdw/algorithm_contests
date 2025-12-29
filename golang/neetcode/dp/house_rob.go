package dp

func rob(nums []int) int {
	a, b := 0, 0
	for _, v := range nums {
		a, b = b, max(a+v, b)
	}
	return b
}
