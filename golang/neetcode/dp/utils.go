package dp

func manacher(ss string) []int {
	if len(ss) == 0 {
		return nil
	}
	bs := []byte{'@', '#'}
	for _, ch := range []byte(ss) {
		bs = append(bs, ch, '#')
	}
	s := string(append(bs, '$'))

	n, md, r := len(s), 0, -1
	disR := make([]int, n)
	for i := 1; i < n-1; i++ {
		d := 0
		if r > i+d {
			mirror := md*2 - i
			d = max(d, min(r-i, disR[mirror]))
		}
		for 0 <= i-d && i+d < n && s[i-d] == s[i+d] {
			d++
		}
		d--
		disR[i] = d
		if i+d >= r {
			r = i + d
			md = i
		}
	}
	return disR[2 : n-1]
}
