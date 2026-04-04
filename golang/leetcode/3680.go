package main

import (
	"math/rand/v2"
)

func generateSchedule(n int) [][]int {
	if n <= 4 {
		return nil
	}
	a := make([][]int, 0, n*(n-1))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a = append(a, []int{i, j}, []int{j, i})
		}
	}
	m := len(a)
	swap := func(i, j int) {
		a[i], a[j] = a[j], a[i]
	}
	valid := func(p, q []int) bool {
		for _, x := range p {
			for _, y := range q {
				if x == y {
					return false
				}
			}
		}
		return true
	}
outer:
	for {
		rand.Shuffle(m, swap)
		for i := 1; i < m; i++ {
			var j int
			for j = i; j < m && (!valid(a[j], a[i-1])); j++ {
			}
			if j >= m {
				continue outer
			}
			swap(i, j)
		}
		return a
	}
}
