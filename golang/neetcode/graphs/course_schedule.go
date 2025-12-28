package graphs

func canFinish(n int, prerequisites [][]int) bool {
	dep := make([][]int, n)
	for _, pre := range prerequisites {
		dep[pre[0]] = append(dep[pre[0]], pre[1])
	}

	vis := make([]bool, n)
	path := make([]bool, n)
	var hasCycle func(i int) bool
	hasCycle = func(i int) bool {
		if vis[i] {
			return path[i]
		}
		vis[i] = true
		path[i] = true
		for _, j := range dep[i] {
			if hasCycle(j) {
				return true
			}
		}
		path[i] = false
		return false
	}

	for i := 0; i < n; i++ {
		if hasCycle(i) {
			return false
		}
	}
	return true
}
