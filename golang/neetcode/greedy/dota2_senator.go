package greedy

func predictPartyVictory1(senate string) string {
	// greedy: ban next-other, as a circular array
	ban := 0 // ban D:+1, ban R: -1
	que := []byte(senate)
	for -len(que) < ban && ban < len(que) {
		c := que[0]
		que = que[1:]
		d := 1
		if c != 'R' {
			d = -1
		}
		if ban*d >= 0 {
			que = append(que, c)
		}
		ban += d
	}
	if que[0] == 'R' {
		return "Radiant"
	}
	return "Dire"
}

func predictPartyVictory(senate string) string {
	// lower impl cognitive burden by convert to more informative: raw que -> maintain each side's indexes
	var r, d []int
	for i, ch := range senate {
		if ch == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	i := len(senate)
	for len(r) > 0 && len(d) > 0 {
		ri, di := r[0], d[0]
		r, d = r[1:], d[1:]
		if ri < di {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
		i++
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}
