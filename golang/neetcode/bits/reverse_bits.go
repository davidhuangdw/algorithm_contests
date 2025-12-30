package bits

func reverseBits(n int) (ans int) {
	for i := 0; i < 32; i++ {
		if n&(1<<i) > 0 {
			ans |= 1 << (31 - i)
		}
	}
	return ans
}
