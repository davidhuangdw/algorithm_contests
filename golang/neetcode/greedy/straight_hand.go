package greedy

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) == 0 {
		return true
	}
	if len(hand)%groupSize != 0 {
		return false
	}
	cnt := make(map[int]int)
	for _, v := range hand {
		cnt[v]++
	}

	for _, v := range hand {
		if cnt[v] == 0 {
			continue
		}
		l := v
		for ; cnt[l-1] > 0; l-- {
		}

		for x := l; x <= v; x++ {
			if cnt[x] == 0 {
				continue
			}
			for y := x + groupSize - 1; y >= x; y-- {
				if cnt[y] < cnt[x] {
					return false
				}
				cnt[y] -= cnt[x]
			}
		}
	}
	return true
}
