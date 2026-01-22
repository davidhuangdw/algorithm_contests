package main

import (
	"fmt"
	"io"
)

func minStr(s string) string {
	n, s := len(s), s+s
	i, j, k := 0, 1, 0
	for i < n && j < n && k < n {
		if s[i+k] == s[j+k] {
			k++
			continue
		}
		if s[i+k] < s[j+k] {
			j += k + 1
		} else {
			i, j = j, i+k+1
		}
		if i == j {
			j++
		}
		k = 0
	}
	return s[i : i+n]
}

func P13270_minStr(IN io.Reader, OUT io.Writer) {
	var s string
	fmt.Fscan(IN, &s)
	fmt.Fscan(IN, &s)
	fmt.Fprintln(OUT, minStr(s))
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P13270_minStr(IN, OUT)
//}
