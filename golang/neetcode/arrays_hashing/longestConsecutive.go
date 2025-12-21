package arrays_hashing

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	mx := 0
	for v := range m {
		if m[v-1] { // check each start-point
			continue
		}
		u := v + 1
		for ; m[u]; u++ {
		}
		mx = max(mx, u-v)
	}
	return mx
}