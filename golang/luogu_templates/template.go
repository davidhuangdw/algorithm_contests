package main

import (
	"bufio"
	"io"
	"os"
)

func P(IN io.Reader, OUT io.Writer) {
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P(IN, OUT)
}
