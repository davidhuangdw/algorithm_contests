package main

import (
	"fmt"
	"io"
)

type BinaryIndexTree1 struct {
	segSum []int
}

func NewBinaryIndexTree1(n int, initArr []int) *BinaryIndexTree1 {
	tr := &BinaryIndexTree1{segSum: make([]int, n+1)}
	if len(initArr) > 0 {
		preSum := make([]int, n+1)
		for i := 1; i <= n; i++ {
			v := initArr[i-1]
			preSum[i] = preSum[i-1] + v
			tr.segSum[i] = preSum[i] - preSum[i&(i-1)]
		}
	}
	return tr
}

func (bi *BinaryIndexTree1) Add(i, v int) {
	for i < len(bi.segSum) {
		bi.segSum[i] += v
		i += i & -i
	}
}
func (bi *BinaryIndexTree1) preSum(i int) int {
	s := 0
	for i > 0 {
		s += bi.segSum[i]
		i &= i - 1
	}
	return s
}
func (bi *BinaryIndexTree1) Sum(l, r int) int {
	return bi.preSum(r) - bi.preSum(l-1)
}

func P3374_BIT1(IN io.Reader, OUT io.Writer) {
	var n, m, op, x, y int
	fmt.Fscan(IN, &n, &m)
	init := make([]int, n)
	for i := range init {
		fmt.Fscan(IN, &init[i])
	}

	tr := NewBinaryIndexTree1(n, init)
	for i := 0; i < m; i++ {
		fmt.Fscan(IN, &op, &x, &y)
		switch op {
		case 1:
			tr.Add(x, y)
		case 2:
			fmt.Fprintln(OUT, tr.Sum(x, y))
		}
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3374_BIT1(IN, OUT)
//}
