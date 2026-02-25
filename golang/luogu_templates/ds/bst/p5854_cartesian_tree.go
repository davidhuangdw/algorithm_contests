package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func buildCartesianTree(w []int) (chd [][2]int) {
	n := len(w) - 1
	chd = make([][2]int, n+1)
	que := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		pre := 0
		for len(que) > 0 && w[que[len(que)-1]] > w[i] {
			pre, que = que[len(que)-1], que[:len(que)-1]
		}
		chd[i][0] = pre
		if len(que) > 0 {
			chd[que[len(que)-1]][1] = i
		}
		que = append(que, i)
	}
	return chd
}

func P5854_cartesian_tree(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	w := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(IN, &w[i])
	}
	chd := buildCartesianTree(w)

	la, ra := 0, 0
	for i := 1; i <= n; i++ {
		la ^= i * (chd[i][0] + 1)
		ra ^= i * (chd[i][1] + 1)
	}
	fmt.Fprintf(OUT, "%v %v\n", la, ra)
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P5854_cartesian_tree(IN, OUT)
}
