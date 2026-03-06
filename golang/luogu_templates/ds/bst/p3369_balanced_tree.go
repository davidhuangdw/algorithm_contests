package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func P3369_balanced_tree(IN io.Reader, OUT io.Writer) {
	var tr MultiSet
	tr = &FhqTreap{}
	var n int
	fmt.Fscan(IN, &n)
	for i := 0; i < n; i++ {
		var op, x int
		fmt.Fscan(IN, &op, &x)
		switch op {
		case 1:
			tr.Insert(x)
		case 2:
			tr.Del(x)
		case 3:
			fmt.Fprintln(OUT, tr.Rank(x))
		case 4:
			fmt.Fprintln(OUT, tr.Kth(x))
		case 5:
			fmt.Fprintln(OUT, tr.Prev(x))
		case 6:
			fmt.Fprintln(OUT, tr.Succ(x))
		}
	}
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P3369_balanced_tree(IN, OUT)
}
