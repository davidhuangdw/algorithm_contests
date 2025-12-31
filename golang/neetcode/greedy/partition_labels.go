package greedy

func partitionLabels(s string) []int {
	ans := make([]int, 0)
	last := make(map[byte]int)
	for i, ch := range []byte(s) {
		last[ch] = i
	}

	for i := 0; i < len(s); {
		r, j := i, i
		for ; j <= r; j++ {
			r = max(r, last[s[j]])
		}
		ans = append(ans, j-i)
		i = j
	}
	return ans
}