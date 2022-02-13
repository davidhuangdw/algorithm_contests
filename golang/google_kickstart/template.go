package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() { ProblemX(os.Stdin, os.Stdout) }

func ProblemX(input io.Reader, output io.Writer) {
	IN := bufio.NewReader(input)
	OUT := bufio.NewWriter(output)
	defer func() {
		OUT.Flush()
	}()
	var T int
	Fscan(IN, &T)
	for icase := 1; icase <= T; icase++ {
		var n, k int
		Fscan(IN, &n, &k)
		Fprintf(OUT, "Case #%v: %v\n", icase, icase)
	}
}
