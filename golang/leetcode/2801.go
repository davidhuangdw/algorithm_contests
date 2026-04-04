package main

func countSteppingNumbers(low string, high string) int {
	n, M := len(high), int(1e9+7)
	bs := make([]byte, n-len(low))
	for i := range bs {
		bs[i] = '0'
	}
	low = string(bs) + low
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	type St struct {
		i, pre       int
		lo, hi, zero bool
	}
	mem := make(map[St]int)
	var count func(i, pre int, lo, hi, zero bool) int
	count = func(i, pre int, lo, hi, zero bool) int {
		if i == n {
			return 1
		}
		st := St{i, pre, lo, hi, zero}
		if _, ok := mem[st]; !ok {
			fr, to := 0, 9
			if lo {
				fr = max(fr, int(low[i]-'0'))
			}
			if hi {
				to = min(to, int(high[i]-'0'))
			}
			res := 0
			for d := fr; d <= to; d++ {
				if zero || abs(d-pre) == 1 {
					res = (res + count(i+1, d, lo && d == fr, hi && d == to, zero && d == 0)) % M
				}
			}
			mem[st] = res
		}
		return mem[st]
	}
	return count(0, -1, true, true, true)
}
