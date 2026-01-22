package main

import (
	"bufio"
	"io"
	"os"
)

func ProblemX(IN io.Reader, OUT io.Writer) {
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	ProblemX(IN, OUT)
}
