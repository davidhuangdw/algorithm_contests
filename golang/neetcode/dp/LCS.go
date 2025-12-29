package dp

func longestCommonSubsequence(a string, b string) int {
	mx := make([]int, len(b)+1)
	for _, ci := range []byte(a) {
		pre := 0
		for j, cj := range []byte(b) {
			old := mx[j+1]
			if ci == cj {
				mx[j+1] = pre + 1
			} else {
				mx[j+1] = max(mx[j+1], mx[j])
			}
			pre = old
		}
	}
	return mx[len(b)]
}
