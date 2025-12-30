package bits

func countBits(n int) []int {
	bits := make([]int, n+1)
	for v := 1; v <= n; v++ {
		bits[v] = bits[v>>1] + (v & 1)
	}
	return bits
}
