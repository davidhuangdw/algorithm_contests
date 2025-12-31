package math_geo

func myPow(x float64, n int) (ans float64) {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	ans = 1
	for n > 0 {
		if n%2 == 1 {
			ans *= x
		}
		n /= 2
		x *= x
	}
	return ans
}
