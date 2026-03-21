package main

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	L := int(1e6)
	area := len(blocked) * (len(blocked) - 1) / 2
	blk := make(map[[2]int]bool)
	for _, b := range blocked {
		blk[[2]int{b[0], b[1]}] = true
	}

	bfs := func(s, t [2]int) bool {
		q, vis := [][2]int{s}, make(map[[2]int]bool)
		vis[s] = true

		for len(q) > 0 && len(vis) <= area {
			si, sj := q[0][0], q[0][1]
			q = q[1:]
			for k, di, dj := 0, 0, 1; k < 4; k++ {
				i, j := si+di, sj+dj
				p := [2]int{i, j}
				if p == t {
					return true
				}
				if 0 <= i && i < L && 0 <= j && j < L && !vis[p] && !blk[p] {
					vis[p] = true
					q = append(q, p)
				}
				di, dj = dj, -di
			}
		}
		return len(vis) > area
	}

	s, t := [2]int{source[0], source[1]}, [2]int{target[0], target[1]}
	return bfs(s, t) && bfs(t, s)
}
