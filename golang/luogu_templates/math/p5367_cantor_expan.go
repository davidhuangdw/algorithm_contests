package main

import (
	"fmt"
	"io"
)

func cantorRank(a []int) int {
	n := len(a)
	MOD := 998244353
	fact, ans := 1, 1

	BIT := make([]int, n+1)
	add := func(v int) {
		for v <= n {
			BIT[v] += 1
			v += v & -v
		}
	}

	pre := func(v int) (sum int) {
		for v > 0 {
			sum += BIT[v]
			v &= v - 1
		}
		return sum
	}

	add(a[n-1])
	for i := 1; i < n; i++ {
		fact = fact * i % MOD
		v := a[n-1-i]
		ans = (ans + pre(v-1)*fact%MOD) % MOD
		add(v)
	}

	return ans
}

func P5367_cantor_expan(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(IN, &a[i])
	}
	fmt.Fprintln(OUT, cantorRank(a))
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P5367_cantor_expan(IN, OUT)
//}
