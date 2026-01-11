package intervals

import (
	"container/heap"
	"slices"
)

func mostBooked1(n int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})
	cnt, ed := make([]int, n), make([]int, n)

	findRoom := func(st int) int {
		minEdRoom := 0
		for i, e := range ed {
			if e <= st {
				return i
			}
			if ed[i] < ed[minEdRoom] {
				minEdRoom = i
			}
		}
		return minEdRoom
	}

	for _, m := range meetings {
		room := findRoom(m[0])
		cnt[room]++
		ed[room] = max(m[1], ed[room]+m[1]-m[0])
	}
	ans := 0
	for i, c := range cnt {
		if c > cnt[ans] {
			ans = i
		}
	}
	return ans
}

func mostBooked(n int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})
	cnt := make([]int, n)

	hp := NewHeap(func(a, b []int) bool {
		if a[0] == b[0] {
			return a[1] <= b[1]
		}
		return a[0] <= b[0]
	})
	for i := 0; i < n; i++ {
		heap.Push(hp, []int{0, i})
	}

	findRoom := func(a int, b int) int {
		for hp.Top()[0] < a {
			r := heap.Pop(hp).([]int)
			heap.Push(hp, []int{a, r[1]})
		}
		r := heap.Pop(hp).([]int)
		heap.Push(hp, []int{r[0] + b - a, r[1]})
		return r[1]
	}

	for _, m := range meetings {
		room := findRoom(m[0], m[1])
		cnt[room]++
	}
	ans := 0
	for i, c := range cnt {
		if c > cnt[ans] {
			ans = i
		}
	}
	return ans
}
