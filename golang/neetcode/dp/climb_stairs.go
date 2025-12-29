package dp

func climbStairs(n int) int {
	a, b := 0, 1 // v for -1, 0
	for ; n > 0; n-- {
		a, b = b, a+b
	}
	return b
}
