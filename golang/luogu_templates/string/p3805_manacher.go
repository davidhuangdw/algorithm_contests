package main

import (
	"fmt"
	"io"
)

func manacher(ss string) []int {
	n := len(ss)*2 + 3
	s := make([]byte, n) // pre-create full length to avoid extend
	s[0], s[n-1] = '@', '$'
	for i := 1; i < n-1; i++ {
		s[i] = '*'
		if i%2 == 0 {
			s[i] = ss[i/2-1]
		}
	}

	dis, md, r := make([]int, n), -1, -1
	for i := 1; i < n-1; i++ {
		j := i + 1
		if i < r {
			j = min(r, i+dis[md*2-i]) + 1
		}
		for s[j] == s[i*2-j] {
			j++
		}
		j--
		if j > r {
			md, r = i, j
		}
		dis[i] = j - i
	}
	return dis[1 : n-2]
}

func P3805_manacher(IN io.Reader, OUT io.Writer) {
	var s string
	fmt.Fscan(IN, &s) // use bufio to read large str
	mx := 0
	for _, d := range manacher(s) {
		mx = max(mx, d)
	}
	fmt.Println(mx)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3805_manacher(IN, OUT)
//}
