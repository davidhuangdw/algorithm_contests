package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func P2197_nim(IN io.Reader, OUT io.Writer) {
	var T int
	fmt.Fscan(IN, &T)
	for ; T > 0; T-- {
		var n int
		fmt.Fscan(IN, &n)
		sm := 0
		for i := 0; i < n; i++ {
			var x int
			fmt.Fscan(IN, &x)
			sm ^= x
		}
		if sm == 0 {
			fmt.Fprintln(OUT, "No")
		} else {
			fmt.Fprintln(OUT, "Yes")
		}
	}
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P2197_nim(IN, OUT)
}
