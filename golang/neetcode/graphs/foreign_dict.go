package graphs

import "slices"

func foreignDictionary(words []string) string {
	es := make(map[byte]map[byte]bool)
	for i, wd := range words {
		for _, ch := range []byte(wd) {
			if _, ok := es[ch]; !ok {
				es[ch] = make(map[byte]bool)
			}
		}
		if i == 0 {
			continue
		}
		pre := []byte(words[i-1])
		n, m, k := len(pre), len(wd), 0
		for ; k < n && k < m && pre[k] == wd[k]; k++ {
		}
		if !(k < n && k < m) {
			if n > m {
				return ""
			}
			continue
		}
		// pre[k] < wd[k]
		es[pre[k]][wd[k]] = true
	}

	var topo []byte
	vis := make(map[byte]int)
	var hasCycle func(u byte) bool
	hasCycle = func(u byte) bool {
		if vis[u] != 0 {
			return vis[u] == 1
		}
		vis[u] = 1                    // 1 ~= in-path
		defer func() { vis[u] = 2 }() // 2 ~= done

		for v, _ := range es[u] {
			if hasCycle(v) {
				return true
			}
		}
		topo = append(topo, u)
		return false
	}
	for u, _ := range es {
		if hasCycle(u) {
			return ""
		}
	}

	slices.Reverse(topo)
	return string(topo)
}
