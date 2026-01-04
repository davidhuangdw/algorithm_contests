package stack

func decodeString(s string) string {
	var ans []byte
	var st [][]int
	num, l := 0, -1
	for _, ch := range []byte(s) {
		switch {
		case '0' <= ch && ch <= '9':
			if num == 0 {
				l = len(ans)
			}
			num = num*10 + int(ch-'0')
		case ch == '[':
			st = append(st, []int{num, l})
			num = 0
		case ch == ']':
			r, p := len(ans), st[len(st)-1]
			st = st[:len(st)-1]
			for k := 0; k < p[0]-1; k++ { // bug: p[0]-1, already have one
				for i := p[1]; i < r; i++ {
					ans = append(ans, ans[i])
				}
			}
		default:
			ans = append(ans, ch)
		}
	}
	return string(ans)
}