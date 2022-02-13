package main

import (
	"bufio"
	. "fmt"
	"io"
)

//func main() { Problem1(os.Stdin, os.Stdout) }

func Problem1(input io.Reader, output io.Writer) {
	IN := bufio.NewReader(input)
	OUT := bufio.NewWriter(output)
	defer func() {
		OUT.Flush()
	}()
	var T int
	Fscan(IN, &T)
	for icase := 1; icase <= T; icase++ {
		var S, F string
		Fscan(IN, &S, &F)
		Fprintf(OUT, "Case #%v: %v\n", icase, MinTransform(S, F))
	}
}

func MinTransform(S, F string) int {
	N := 26
	dis := make([]int, N)
	for i := range dis {
		dis[i] = N
	}
	for _, ch := range F {
		i := int(ch - 'a')
		for d := 0; d < N; d++ {
			j := (i + d) % N
			update := func() {
				if d < dis[j] {
					dis[j] = d
				}
			}
			update()
			j = (i - d + N) % N
			update()
		}
	}

	sum := 0
	for _, ch := range S {
		sum += dis[ch-'a']
	}
	return sum
}
