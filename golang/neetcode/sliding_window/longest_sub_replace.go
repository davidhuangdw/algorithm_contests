package sliding_window

func characterReplacement1(s string, k int) int { // O(n) by iter each-char's pos-arr
	ss := []byte(s)
	chars := make(map[byte][]int)
	for i, c := range ss {
		chars[c] = append(chars[c], i)
	}

	mx := 0
	for _, pos := range chars { // check each char
		i, acc := 0, 0
		for _, r := range pos {
			acc++
			for r+1-pos[i]-acc > k { // holes > k
				i++
				acc--
			}
			mx = max(mx, acc+k)
		}
	}
	return min(mx, len(ss))
}

func characterReplacement(s string, k int) int { // by iter valid substr
	ss := []byte(s)
	maxF, i, cnt := 0, 0, make(map[byte]int)
	for j, c := range ss {
		cnt[c]++
		maxF = max(maxF, cnt[c])

		if j+1-i > maxF+k {
			cnt[ss[i]]--
			i++
		}
	}
	return min(maxF+k, len(ss))
}
