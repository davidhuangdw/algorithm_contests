package greedy

func lemonadeChange(bills []int) bool {
	val := []int{20, 10, 5}
	cnt := make(map[int]int)
	for _, b := range bills {
		back := b - 5
		i := 0
		for back > 0 && i < len(val) {
			if cnt[val[i]] > 0 && val[i] <= back {
				back -= val[i]
				cnt[val[i]]--
			} else {
				i++
			}
		}
		if back > 0 {
			return false
		}
		cnt[b]++
	}
	return true
}
