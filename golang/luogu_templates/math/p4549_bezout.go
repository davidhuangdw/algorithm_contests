package main

import (
	"fmt"
	"io"
)

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func P4549_bezout(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	g := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(IN, &x)
		g = gcd(g, x)
	}
	fmt.Fprintln(OUT, g)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P4549_bezout(IN, OUT)
//}
