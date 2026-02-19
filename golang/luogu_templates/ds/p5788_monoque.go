package main

import (
	"fmt"
	"io"
)

func nextLarge(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		v := a[i]
		j := i + 1
		for 0 < j && a[j] <= v {
			j = ans[j] - 1
		}
		ans[i] = j + 1
	}
	return ans
}

func P5788_monoque(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &a[i])
	}
	ans := nextLarge(a)
	for i, v := range ans {
		if i > 0 {
			fmt.Fprint(OUT, " ")
		}
		fmt.Fprint(OUT, v)
	}
	fmt.Fprintln(OUT)
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P5788_monoque(IN, OUT)
//}
