package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"strings"
)

//func main() { Problem3(os.Stdin, os.Stdout) }

func Problem3(input io.Reader, output io.Writer) {
	IN := bufio.NewReader(input)
	OUT := bufio.NewWriter(output)
	defer func() {
		OUT.Flush()
	}()
	var T int
	Fscan(IN, &T)
	for icase := 1; icase <= T; icase++ {
		var n int
		Fscan(IN, &n)
		C := make([]int, n)
		for i := range C {
			Fscan(IN, &C[i])
		}
		Fprintf(OUT, "Case #%v: %v\n", icase, hIndex(C))
	}
}

type Que []int

func (q Que) Len() int           { return len(q) }
func (q Que) Less(i, j int) bool { return q[i] < q[j] }
func (q Que) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *Que) Push(x interface{}) { *q = append(*q, x.(int)) }
func (q *Que) Pop() interface{} {
	n := q.Len()
	x := (*q)[n-1]
	*q = (*q)[:n-1]
	return x
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func hIndex(C []int) string {
	var b strings.Builder
	var que Que
	for i, x := range C {
		if x > len(que) {
			heap.Push(&que, x)
		}
		if que[0] < len(que) {
			heap.Pop(&que)
		}
		//Printf("%#v\n", que)

		if i != 0 {
			b.WriteRune(' ')
		}
		b.WriteString(Sprintf("%v", len(que)))
	}
	return b.String()
}
