package greedy

func mergeTriplets(triplets [][]int, target []int) bool {
	mx := make([]int, 3)
	for _, t := range triplets {
		ok := true
		for i := 0; ok && i < 3; i++ {
			ok = t[i] <= target[i]
		}
		if !ok {
			continue
		}
		for i := 0; i < 3; i++ {
			mx[i] = max(mx[i], t[i])
		}
	}
	for i := 0; i < 3; i++ {
		if mx[i] != target[i] {
			return false
		}
	}
	return true
}
