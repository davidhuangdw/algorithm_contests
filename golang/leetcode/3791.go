package main

func countBalanced(low int64, high int64) int64 {
	D := 1
	for x := high; x > 0; x /= 10 {
		D++
	}
	dlow, dhigh := make([]int, D), make([]int, D)
	for i, b := 0, 1; i < D; i++ {
		dlow[i] = int(low) / b % 10
		dhigh[i] = int(high) / b % 10
		b *= 10
	}

	type St struct {
		i, diff   int
		low, high bool
	}
	mem := make(map[St]int)
	var dp func(i, diff int, low, high bool) int // time-complexity: only need diff instead of sum-pair
	dp = func(i, diff int, low, high bool) int {
		if i < 0 {
			if diff == 0 {
				return 1
			}
			return 0
		}
		s := St{i, diff, low, high}
		if _, ok := mem[s]; !ok {
			fr, to := 0, 9
			if low {
				fr = dlow[i]
			}
			if high {
				to = dhigh[i]
			}
			res, mul := 0, 1
			if i&1 > 0 {
				mul = -1
			}
			for d := fr; d <= to; d++ {
				res += dp(i-1, diff+mul*d, low && d == dlow[i], high && d == dhigh[i])
			}
			mem[s] = res
		}
		return mem[s]
	}
	return int64(dp(D-1, 0, true, true))
}
