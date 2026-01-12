package bits

import "slices"

func addBinary(a string, b string) string {
	c, n, m := 0, len(a), len(b)
	var res []byte
	for i := 0; i < n || i < m || c > 0; i++ {
		if i < n {
			c += int(a[n-1-i] - '0')
		}
		if i < m {
			c += int(b[m-1-i] - '0')
		}
		res = append(res, byte('0'+(c&1)))
		c >>= 1
	}
	slices.Reverse(res)
	return string(res)
}
