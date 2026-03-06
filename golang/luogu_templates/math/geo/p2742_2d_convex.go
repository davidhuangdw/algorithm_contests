package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

const eps = 1e-8

func convex2d(points [][2]float64) [][2]float64 {
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		return a[0] < b[0]-eps || (a[0] < b[0]+eps && a[1]+eps < b[1])
	})
	slopInc := func(a, b, c [2]float64) bool { // bc's slop > ab's slop
		v1 := [2]float64{b[0] - a[0], b[1] - a[1]}
		v2 := [2]float64{c[0] - b[0], c[1] - b[1]}
		return v1[0]*v2[1]-v1[1]*v2[0] > 0 // cross-product > 0
	}
	n := len(points)
	stk := make([]int, 0, n+1)
	for i := 0; i < n; i++ {
		l := len(stk)
		for l >= 2 && !slopInc(points[stk[l-2]], points[stk[l-1]], points[i]) {
			stk = stk[:len(stk)-1]
			l = len(stk)
		}
		stk = append(stk, i)
	}
	below := len(stk)
	for i := n - 1; i >= 0; i-- {
		l := len(stk)
		for l > below && !slopInc(points[stk[l-2]], points[stk[l-1]], points[i]) {
			stk = stk[:len(stk)-1]
			l = len(stk)
		}
		stk = append(stk, i)
	}
	stk = stk[:len(stk)-1]
	circle := make([][2]float64, len(stk))
	for i, pi := range stk {
		circle[i] = points[pi]
	}
	//fmt.Printf("------------%v:%#v:%#v", below, stk, used)
	//fmt.Println("------------")
	//for i := 0; i < below; i++ {
	//	fmt.Println(circle[i][0], " ", circle[i][1])
	//}
	//fmt.Println("------------")
	//for i := below; i < len(circle); i++ {
	//	fmt.Println(circle[i][0], " ", circle[i][1])
	//}
	return circle
}

func P2742_2d_convex(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	points := make([][2]float64, n)
	for i := range points {
		fmt.Fscan(IN, &points[i][0], &points[i][1])
	}

	circle := convex2d(points)
	dis := func(a, b [2]float64) float64 {
		dx, dy := a[0]-b[0], a[1]-b[1]
		return math.Sqrt(dx*dx + dy*dy)
	}
	var sum float64
	for i := 0; i < len(circle); i++ {
		sum += dis(circle[i], circle[(i+1)%len(circle)])
	}
	fmt.Fprintf(OUT, "%.2f\n", sum)
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P2742_2d_convex(IN, OUT)
}
