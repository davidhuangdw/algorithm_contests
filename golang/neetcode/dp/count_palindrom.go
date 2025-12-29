package dp

func countSubstrings(s string) int {
	sum := 0
	for _, l := range manacher(s) {
		sum += (l + 1) / 2
	}
	return sum
}
