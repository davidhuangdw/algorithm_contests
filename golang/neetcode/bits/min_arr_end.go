package bits

func minEnd(n int, x int) int {
	rem, bits, ans := n-1, 0, x
	for y := x; y > 0 && rem > 0; y >>= 1 {
		if y&1 == 0 {
			ans += (rem & 1) << bits
			rem >>= 1
		}
		bits++
	}
	ans += rem << bits
	return ans
}
