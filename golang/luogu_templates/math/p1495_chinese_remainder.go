package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func exgcd(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := exgcd(b, a%b) // b*x1 + a%b*y1 == (a/b*b+a%b)*y1 + b*(x1-a/b*y1)
	return g, y1, x1 - a/b*y1
}

func chineseRemainder(ab [][2]int) int {
	mul := 1
	for _, p := range ab {
		mul *= p[0]
	}

	ans := 0
	for _, p := range ab {
		a, b := p[0], p[1]
		oth := mul / a           // oth % other-a == 0
		_, x, _ := exgcd(oth, a) // (oth * x) % a = 1 => (oth * x * b) % a  == b
		ans = (ans + oth*x*b%mul + mul) % mul
	}
	return ans
}

func P1495_chinese_remainder(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	ab := make([][2]int, n)
	for i := range ab {
		var a, b int
		fmt.Fscan(IN, &a, &b)
		ab[i] = [2]int{a, b}
	}
	fmt.Fprintln(OUT, chineseRemainder(ab))
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P1495_chinese_remainder(IN, OUT)
}
