package math_geo

func spiralOrder1(matrix [][]int) []int {
	h, w := len(matrix), len(matrix[0])
	ans := make([]int, 0, h*w)
	i, j, d := 0, -1, 1
	for h > 0 && w > 0 {
		for k := 0; k < w; k++ {
			j += d
			ans = append(ans, matrix[i][j])

		}
		for k := 0; k < h-1; k++ {
			i += d
			ans = append(ans, matrix[i][j])
		}
		h--
		w--
		d = -d
	}
	return ans
}

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	ans, i, j := make([]int, 0, n*m), 0, -1
	for dis, vertDis, di, dj := m, n, 0, 1; dis > 0; dis, vertDis, di, dj = vertDis-1, dis, dj, -di {
		for k := 0; k < dis; k++ {
			i, j = i+di, j+dj
			ans = append(ans, matrix[i][j])
		}
	}
	return ans
}
