package main

import (
	"fmt"
	"io"
)

type BinaryIndexTree2 struct {
	segAdd []int
}

func NewBinaryIndexTree2(n int, initArr []int) *BinaryIndexTree2 {
	tr := &BinaryIndexTree2{segAdd: make([]int, n+1)}
	if len(initArr) > 0 {
		for i := 1; i <= n; i++ {
			parent := i + (i & -i)
			// every segAdd[] store reverse-diff := v[i]-v[father]
			tr.segAdd[i] = initArr[i-1]
			if parent <= n {
				tr.segAdd[i] -= initArr[parent-1]
			}
		}
	}
	return tr
}

func (bi *BinaryIndexTree2) Get(i int) int {
	s := 0
	for i < len(bi.segAdd) {
		s += bi.segAdd[i]
		i += i & -i
	}
	return s
}
func (bi *BinaryIndexTree2) preAdd(i, v int) {
	for i > 0 {
		bi.segAdd[i] += v
		i &= i - 1
	}
}
func (bi *BinaryIndexTree2) Add(l, r, v int) {
	bi.preAdd(r, v)
	bi.preAdd(l-1, -v)
}

func P3368_BIT2(IN io.Reader, OUT io.Writer) {
	var n, m, op, x, y, v int
	fmt.Fscan(IN, &n, &m)
	init := make([]int, n)
	for i := range init {
		fmt.Fscan(IN, &init[i])
	}
	fmt.Fscanln(IN)

	tr := NewBinaryIndexTree2(n, init)
	for i := 0; i < m; i++ {
		fmt.Fscanln(IN, &op, &x, &y, &v)
		switch op {
		case 1:
			tr.Add(x, y, v)
		case 2:
			fmt.Fprintln(OUT, tr.Get(x))
		}
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3368_BIT2(IN, OUT)
//}
