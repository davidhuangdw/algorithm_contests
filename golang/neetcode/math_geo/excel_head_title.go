package math_geo

func convertToTitle(columnNumber int) string {
	var b []byte
	x := columnNumber
	for x > 0 {
		b = append(b, byte('A'+(x-1)%26))
		x = (x - 1) / 26
	}

	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	return string(b)
}
