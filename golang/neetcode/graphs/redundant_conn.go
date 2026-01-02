package graphs

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	pa := make([]int, n+1)
	rank := make([]int, n+1)
	for i := 0; i < n; i++ {
		pa[i] = i
		rank[i] = 1
	}
	var find func(i int) int
	find = func(i int) int {
		if pa[i] != i {
			pa[i] = find(pa[i])
		}
		return pa[i]
	}
	union := func(i int, j int) bool {
		i, j = find(i), find(j)
		if i == j {
			return false
		}
		if rank[i] < rank[j] {
			i, j = j, i
		}
		pa[j] = i
		rank[i] += rank[j]
		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}
