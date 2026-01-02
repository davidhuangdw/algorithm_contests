package graphs

import "slices"

func findItinerary(tickets [][]string) []string {
	nxt := make(map[string][]string)
	for _, t := range tickets {
		u, v := t[0], t[1]
		nxt[u] = append(nxt[u], v)
	}
	for u, _ := range nxt {
		slices.Sort(nxt[u])
	}

	var path []string
	var dfs func(u string)
	dfs = func(u string) {
		for len(nxt[u]) > 0 {
			v := nxt[u][0]
			nxt[u] = nxt[u][1:]
			dfs(v)
		}
		path = append(path, u)
	}
	dfs("JFK")
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
