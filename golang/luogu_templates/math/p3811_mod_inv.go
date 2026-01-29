package main

import (
	"fmt"
	"io"
)

func modInv(n, p int) []int {
	inv := make([]int, n+1)
	inv[1] = 1
	for v := 2; v <= n; v++ {
		inv[v] = p - p/v*inv[p%v]%p
	}
	return inv
}

func P3811_mod_inv(IN io.Reader, OUT io.Writer) {
	var n, p int
	fmt.Fscan(IN, &n, &p)
	inv := modInv(n, p)
	for _, x := range inv[1:] {
		fmt.Fprintln(OUT, x)
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3811_mod_inv(IN, OUT)
//}
