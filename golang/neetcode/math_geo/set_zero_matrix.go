package math_geo

func setZeroes(matrix [][]int) { // O(1) inplace set
	n, m := len(matrix), len(matrix[0])
	r0, c0 := 1, 1
	// cache info of 1st-row/col
	for i := 0; c0 == 1 && i < n; i++ {
		if matrix[i][0] == 0 {
			c0 = 0
		}
	}
	for j := 0; r0 == 1 && j < m; j++ {
		if matrix[0][j] == 0 {
			r0 = 0
		}
	}

	// store onto 1st-row/col
	for i, row := range matrix {
		for j, _ := range row {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// update matrix by 1st-row/col
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// handle 1st-row/col
	if c0 == 0 {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}

	if r0 == 0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}

}
