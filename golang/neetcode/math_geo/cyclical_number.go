package math_geo

func isHappy1(n int) bool {
	exist := make(map[int]bool)
	for n > 1 {
		if exist[n] {
			return false
		}
		exist[n] = true
		m := 0
		for n > 0 {
			m += (n % 10) * (n % 10)
			n /= 10
		}
		n = m
	}
	return n == 1
}

func isHappy2(n int) bool {
	nxt := func(v int) (s int) {
		for v > 0 {
			s += (v % 10) * (v % 10)
			v /= 10
		}
		return s
	}
	slow, fast := n, nxt(n)
	for slow != fast {
		slow = nxt(slow)
		fast = nxt(nxt(fast))
	}
	return fast == 1
}

func isHappy(n int) bool { // Brent’s Algorithm
	nxt := func(v int) (s int) {
		for v > 0 {
			s += (v % 10) * (v % 10)
			v /= 10
		}
		return s
	}
	slow, fast := n, nxt(n)
	limit, step := 1, 0
	for slow != fast {
		if step > limit {
			slow = fast
			limit *= 2
			step = 0
		}
		fast = nxt(fast)
		step++
	}
	return fast == 1
}
