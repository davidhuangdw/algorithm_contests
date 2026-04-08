package main

func getProbability(balls []int) float64 {
	n, sum := len(balls), 0
	for _, v := range balls {
		sum += v
	}
	hf := sum / 2
	comb := func(n, k int) float64 {
		res := 1
		for i := 1; i <= k; i++ {
			res = res * (n + 1 - i) / i
		}
		return float64(res)
	}
	mm := make(map[[3]int]float64)
	var dfs func(i, dcolor, ln int) float64
	dfs = func(i, dcolor, ln int) float64 {
		if i == n {
			if dcolor == 0 && ln == hf {
				return 1
			}
			return 0
		}
		st := [3]int{i, dcolor, ln}
		if _, ok := mm[st]; !ok {
			b, res := balls[i], .0
			for l := 0; l <= b; l++ {
				dc := dcolor
				switch l {
				case 0:
					dc++
				case b:
					dc--
				}
				res += dfs(i+1, dc, ln+l) * comb(b, l)
			}
			mm[st] = res
		}
		return mm[st]
	}
	return float64(dfs(0, 0, 0)) / comb(sum, hf)
}
