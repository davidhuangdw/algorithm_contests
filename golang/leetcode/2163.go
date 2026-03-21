package main

import "container/heap"

func minimumDifference(nums []int) int64 {
	n := len(nums) / 3
	div := make([]int, n+1)

	sum := 0
	left := NewHeap(func(a, b int) bool {
		return a > b
	})
	for i := 0; i < 2*n; i++ {
		v := nums[i]
		heap.Push(left, v)
		sum += v
		if left.Len() > n {
			sum -= heap.Pop(left).(int)
		}
		if i >= n-1 {
			div[i-(n-1)] += sum
		}
	}

	sum = 0
	right := NewHeap(func(a, b int) bool {
		return a < b
	})
	for i := n*3 - 1; i >= n; i-- {
		v := nums[i]
		heap.Push(right, v)
		sum += v
		if right.Len() > n {
			sum -= heap.Pop(right).(int)
		}
		if i-n <= n {
			div[i-n] -= sum
		}
	}

	ans := div[0]
	for _, v := range div {
		ans = min(ans, v)
	}
	return int64(ans)
}
