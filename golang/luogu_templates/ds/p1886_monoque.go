package main

import (
	"fmt"
	"io"
)

func windowMin(nums []int, k int, less func(a, b int) bool) []int {
	que := make([][2]int, 0, len(nums))
	ans := make([]int, len(nums))
	for i, v := range nums {
		for len(que) > 0 && !less(que[len(que)-1][0], v) {
			que = que[:len(que)-1]
		}
		que = append(que, [2]int{v, i})
		ans[i] = que[0][0]
		if i-que[0][1] >= k-1 {
			que = que[1:]
		}
	}
	return ans
}

func P1886_monoque(IN io.Reader, OUT io.Writer) {
	var n, k int
	fmt.Fscan(IN, &n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &arr[i])
	}

	for _, less := range []func(a, b int) bool{
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a > b },
	} {
		for i, v := range windowMin(arr, k, less)[k-1:] {
			if i > 0 {
				fmt.Fprint(OUT, " ")
			}
			fmt.Fprint(OUT, v)
		}
		fmt.Fprintln(OUT)
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P1886_monoque(IN, OUT)
//}
