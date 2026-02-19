package main

import (
	"fmt"
	"io"
	"slices"
)

func discretize(nums []int) map[int]int {
	a, n := append([]int{}, nums...), len(nums)
	slices.Sort(a)
	rank := make(map[int]int)
	rank[a[0]] = 1
	for i := 1; i < n; i++ {
		if a[i] != a[i-1] {
			rank[a[i]] = rank[a[i-1]] + 1
		}
	}
	return rank
}

func B3694_discretize(IN io.Reader, OUT io.Writer) {
	var T, n int
	fmt.Fscan(IN, &T)
	for ; T > 0; T-- {
		fmt.Fscan(IN, &n)
		nums := make([]int, n)
		for i := range nums {
			fmt.Fscan(IN, &nums[i])
		}

		rank := discretize(nums)

		for i, v := range nums {
			if i > 0 {
				fmt.Fprint(OUT, " ")
			}
			fmt.Fprint(OUT, rank[v])
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
//	B3694_discretize(IN, OUT)
//}
