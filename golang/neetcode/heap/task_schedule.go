package heap

func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}
	var cnt [26]int
	for _, t := range tasks {
		cnt[t-'A']++
	}
	mx := 0
	for _, c := range cnt {
		mx = max(mx, c)
	}
	k := 0
	for _, c := range cnt {
		if c == mx {
			k++
		}
	}
	return max(mx+(mx-1)*n+k-1, len(tasks))
}
