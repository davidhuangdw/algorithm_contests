package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count2DPoints(nums []int, queries [][3]int) []int {
	n, m, mx := len(nums)-1, len(queries), 0
	for _, v := range nums {
		mx = max(mx, v)
	}
	ans, add := make([]int, m), make([][][3]int, n+1)
	for i, q := range queries {
		l, r, x := q[0], q[1], min(q[2], mx)
		if l-1 > 0 {
			add[l-1] = append(add[l-1], [3]int{i, x, -1})
		}
		add[r] = append(add[r], [3]int{i, x, 1})
	}

	BIT := make([]int, mx+1)
	insert := func(x int) {
		for x <= mx {
			BIT[x] += 1
			x += x & -x
		}
	}
	preSum := func(x int) (sum int) {
		for x > 0 {
			sum += BIT[x]
			x &= x - 1
		}
		return sum
	}
	for i := 1; i <= n; i++ {
		x := nums[i]
		insert(x)
		for _, q := range add[i] {
			qi, qx, flag := q[0], q[1], q[2]
			ans[qi] += preSum(qx) * flag
		}
	}
	return ans
}

func P10814_discrete_two_dims(IN io.Reader, OUT io.Writer) {
	var n, m int
	fmt.Fscan(IN, &n, &m)
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(IN, &nums[i])
	}
	queries := make([][3]int, m)
	for i := range queries {
		var l, r, x int
		fmt.Fscan(IN, &l, &r, &x)
		queries[i] = [3]int{l, r, x}
	}
	ans := count2DPoints(nums, queries)
	for _, v := range ans {
		fmt.Fprintln(OUT, v)
	}
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer func() {
		OUT.Flush()
	}()
	P10814_discrete_two_dims(IN, OUT)
}
