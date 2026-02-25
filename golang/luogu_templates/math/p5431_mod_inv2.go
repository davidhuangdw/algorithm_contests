package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func modInv2(a, p int) int {
	n := p - 2
	b, ans := a, 1
	for n > 0 {
		if n&1 > 0 {
			ans = ans * b % p
		}
		b = b * b % p
		n >>= 1
	}
	return ans
}

func p5431(a []int, n, p, k int) int {
	ks, ans := k, 0
	suf := make([]int, n+1)
	suf[n] = 1
	for i := n - 1; i >= 0; i-- {
		suf[i] = a[i] * suf[i+1] % p
	}
	invAll := modInv2(suf[0], p)
	pre := 1
	for i, v := range a {
		inv := (invAll * pre % p) * suf[i+1] % p
		ans = (ans + inv*ks) % p
		pre = pre * v % p
		ks = ks * k % p
	}
	return ans
}

func P5431_mod_inv2(IN io.Reader, OUT io.Writer) {
	var n, p, k int
	fmt.Fscan(IN, &n, &p, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(IN, &a[i])
	}
	fmt.Fprintln(OUT, p5431(a, n, p, k))
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P5431_mod_inv2(IN, OUT)
}
