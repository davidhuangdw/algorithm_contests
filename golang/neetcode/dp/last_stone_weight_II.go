package dp

func lastStoneWeightII(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	// a+b == sum -> a-b=sum-2*b -> min(a-b)==max(b <= sum/2)
	sum := 0
	for _, v := range stones {
		sum += v
	}
	half := sum / 2
	reach, mx := make(map[int]bool), 0
	reach[0] = true
	for _, v := range stones {
		nxt := make(map[int]bool)
		for x := range reach {
			nxt[x] = true
			if x+v <= half {
				nxt[x+v] = true
				mx = max(mx, x+v)
			}
		}
		reach = nxt
	}

	return sum - mx*2
}

func lastStoneWeightII1(stones []int) int { // by bit, if small-range: sum/2 <= 64
	dp, sum := 1, 0
	for _, v := range stones {
		sum += v
		dp |= dp << v
	}
	half := sum / 2
	for v := half; v >= 0; v-- {
		if dp&(1<<v) > 0 {
			return sum - v*2
		}
	}
	return -1
}
