package binary_search

func searchMatrix1(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}

	l, r := 0, n-1 // <=, >
	for l <= r {
		md := (l + r) / 2
		if matrix[md][0] <= target {
			l = md + 1
		} else {
			r = md - 1
		}
	}
	if r < 0 {
		return false
	}

	row := matrix[r]
	l, r = 0, len(row)-1
	for l <= r {
		md := (l + r) / 2
		if row[md] == target {
			return true
		}
		if row[md] < target {
			l = md + 1
		} else {
			r = md - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool { // one pass by flatten
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	l, r := 0, n*m-1
	for l <= r {
		md := (l + r) / 2
		v := matrix[md/m][md%m]
		if v == target {
			return true
		}
		if v < target {
			l = md + 1
		} else {
			r = md - 1
		}
	}
	return false
}
