package math_geo

type CountSquares struct {
	xy map[int]map[int]int
}

func Constructor() CountSquares {
	return CountSquares{xy: make(map[int]map[int]int)}
}

func (co *CountSquares) Add(point []int) {
	x, y := point[0], point[1]
	if _, ok := co.xy[x]; !ok {
		co.xy[x] = make(map[int]int)
	}
	co.xy[x][y]++
}

func (co *CountSquares) Count(point []int) int {
	ans, x, y := 0, point[0], point[1]
	for y1, cnt := range co.xy[x] {
		if y1 == y {
			continue
		}
		d := y1 - y
		for _, x1 := range []int{x - d, x + d} {
			ans += cnt * co.cnt(x1, y) * co.cnt(x1, y1)
		}
	}
	return ans
}
func (co *CountSquares) cnt(x, y int) int {
	if co.xy[x] == nil {
		return 0
	}
	return co.xy[x][y]
}
