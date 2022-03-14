package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

//func main() { Problem2(os.Stdin, os.Stdout) }

func Problem2(input io.Reader, output io.Writer) {
	IN := bufio.NewReader(input)
	OUT := bufio.NewWriter(output)
	defer func() {
		OUT.Flush()
	}()
	var T int
	Fscan(IN, &T)
	for icase := 1; icase <= T; icase++ {
		var K string
		Fscan(IN, &K)
		Fprintf(OUT, "Case #%v: %v\n", icase, rule(K))
	}
}

const VOWEL = "AEIOUaeiou"

func rule(K string) string {
	var owner string
	ch := rune(K[len(K)-1])
	switch {
	case ch == 'y' || ch == 'Y':
		owner = "nobody"
	case strings.ContainsRune(VOWEL, ch):
		owner = "Alice"
	default:
		owner = "Bob"
	}
	return Sprintf("%v is ruled by %v.", K, owner)
}
