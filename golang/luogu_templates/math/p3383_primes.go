package main

import (
	"fmt"
	"io"
)

func getPrimes(n int) []int {
	n = min(n, 1e8)
	ps := make([]int, 0)
	notPri := make([]bool, n+1)
	for v := 2; v <= n; v++ {
		if !notPri[v] {
			ps = append(ps, v)
		}
		for _, p := range ps {
			if p*v > n {
				break
			}
			notPri[p*v] = true
			if v%p == 0 {
				break
			}
		}
	}
	return ps
}

func P3383_primes(IN io.Reader, OUT io.Writer) {
	var n, q, k int
	fmt.Fscan(IN, &n)
	ps := getPrimes(n)

	fmt.Fscan(IN, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(IN, &k)
		fmt.Fprintln(OUT, ps[k-1])
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3383_primes(IN, OUT)
//}
