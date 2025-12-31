package greedy

func checkValidString1(s string) bool {
	var a, b []int
	for i, ch := range s {
		switch ch {
		case '(':
			a = append(a, i)
		case '*':
			b = append(b, i)
		case ')':
			if len(a) > 0 {
				a = a[:len(a)-1] // consume '(' first: use right-most
			} else if len(b) > 0 {
				b = b[1:] // consume left-most star
			} else {
				return false
			}
		}
	}
	for len(a) > 0 {
		if len(b) == 0 || b[len(b)-1] < a[len(a)-1] {
			return false
		}
		a = a[:len(a)-1]
		b = b[:len(b)-1]
	}
	return true
}

func checkValidString(s string) bool {
	avail, unMatchL := 0, 0
	for _, ch := range s {
		switch ch {
		case '(':
			avail++
			unMatchL++
		case '*':
			avail++
			unMatchL = max(0, unMatchL-1)
		case ')':
			if avail <= 0 {
				return false
			}
			avail--
			unMatchL = max(0, unMatchL-1)
		}
	}
	return unMatchL == 0
}
