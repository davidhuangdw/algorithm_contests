package math_geo

func multiply(a string, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}
	n, m := len(a), len(b)
	mul := make([]int, n+m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			s := mul[i+j] + int((a[n-1-i]-'0')*(b[m-1-j]-'0'))
			mul[i+j] = s % 10
			mul[i+j+1] += s / 10
		}
	}
	h := n + m - 2
	if mul[h] > 10 {
		mul[h], mul[h+1] = mul[h]%10, mul[h]/10
	}
	if mul[h+1] == 0 {
		mul = mul[:h+1]
	}

	n, bs := len(mul), make([]byte, 0)
	for i := 0; i < n; i++ {
		bs = append(bs, byte('0'+mul[n-1-i]))
	}
	return string(bs)
}
