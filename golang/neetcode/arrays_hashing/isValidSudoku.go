package arrays_hashing

func isValidSudoku(board [][]byte) bool {
	valid := func(unit []byte) bool {
		var vis [10]bool
		for i := 0; i < 9; i++ {
			if unit[i] == '.' {
				continue
			}
			v := unit[i] - '0'
			if !(1 <= v && v <= 9 && !vis[v]) {
				return false
			}
			vis[v] = true
		}
		return true
	}

	tmp := make([]byte, 9)
	for i := 0; i < 9; i++ {
		// row
		if !valid(board[i]) {
			return false
		}

		// col
		for j := 0; j < 9; j++ {
			tmp[j] = board[j][i]
		}
		if !valid(tmp) {
			return false
		}

		//block
		si, sj := i/3*3, i%3*3 // bug: *3
		for j := 0; j < 9; j++ {
			tmp[j] = board[si+j/3][sj+j%3]
		}
		if !valid(tmp) {
			return false
		}
	}
	return true
}
