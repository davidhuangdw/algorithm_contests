package main

func kSimilarity(s1 string, s2 string) int {
	var a, b []byte
	for i := range []byte(s1) {
		if s1[i] != s2[i] {
			a = append(a, s1[i])
			b = append(b, s2[i])
		}
	}
	if len(a) == 0 {
		return 0
	}
	sa, sb := string(a), string(b)
	qs, qd, vis := []string{sa}, []int{0}, map[string]bool{sa: true}

	for len(qs) > 0 {
		s, d := []byte(qs[0]), qd[0]+1
		qs, qd = qs[1:], qd[1:]
		var di int // diff pos
		for di = 0; s[di] == sb[di]; di++ {
		}
		for i := di + 1; i < len(s); i++ {
			if s[i] == sb[di] {
				s[di], s[i] = s[i], s[di]
				nxt := string(s)
				if !vis[nxt] { // prune
					vis[nxt] = true
					if nxt == sb {
						return d
					}
					qs, qd = append(qs, nxt), append(qd, d)
				}
				s[di], s[i] = s[i], s[di]
			}
		}
	}
	return -1
}
