package main

func numberOfEdgesAdded(n int, edges [][]int) int {
	fa, ws := make([]int, n), make([]int, n)
	for i := range fa {
		fa[i] = i
	}

	var find func(int) (int, int)
	find = func(i int) (int, int) {
		if fa[i] != i {
			r, w := find(fa[i])
			fa[i] = r
			ws[i] ^= w
		}
		return fa[i], ws[i]
	}

	ans := 0
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		ru, wu := find(u)
		rv, wv := find(v)

		if ru == rv {
			if wu^wv == w {
				ans++
			}
		} else {
			fa[rv] = ru
			ws[rv] = wu ^ wv ^ w
			ans++
		}
	}
	return ans
}
