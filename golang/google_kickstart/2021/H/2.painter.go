package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

var (
	Red    = 1
	Yellow = 1 << 1
	Blue   = 1 << 2
	Colors map[rune]int
)

func init() {
	Colors = map[rune]int{
		'U': 0,
		'R': Red,
		'Y': Yellow,
		'B': Blue,
		'O': Red + Yellow,
		'P': Red + Blue,
		'G': Blue + Yellow,
		'A': Red + Blue + Yellow,
	}
}

func main() { Problem2(os.Stdin, os.Stdout) }

func Problem2(input io.Reader, output io.Writer) {
	IN := bufio.NewReader(input)
	OUT := bufio.NewWriter(output)
	defer func() {
		OUT.Flush()
	}()
	var T int
	Fscan(IN, &T)
	for icase := 1; icase <= T; icase++ {
		var N int
		var P string
		Fscan(IN, &N, &P)
		Fprintf(OUT, "Case #%v: %v\n", icase, MinPaints(P))
	}
}

func MinPaints(P string) int {
	sum := 0
	for _, base := range []int{Red, Yellow, Blue} {
		pre := false
		for _, ch := range P {
			if base&Colors[ch] > 0 {
				if !pre {
					sum++
				}
				pre = true
			} else {
				pre = false
			}
		}
	}
	return sum
}
