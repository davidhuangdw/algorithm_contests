package math_geo

func rotate(matrix [][]int) {
	n := len(matrix)
	for k := 0; k < n/2; k++ {
		for i, j := k, k; j <= n-2-k; j++ {
			pre := matrix[i][j]
			for u := 0; u < 4; u++ {
				i, j = j, n-1-i
				matrix[i][j], pre = pre, matrix[i][j]
			}
		}
	}
}
