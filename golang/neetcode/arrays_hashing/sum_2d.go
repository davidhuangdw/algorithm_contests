package arrays_hashing

type NumMatrix struct {
	pre [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	pre := make([][]int, len(matrix)+1)
	pre[0] = make([]int, len(matrix[0])+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, len(matrix[0])+1)
		for j, v := range row {
			pre[i+1][j+1] = v + pre[i+1][j] + pre[i][j+1] - pre[i][j]
		}
	}
	return NumMatrix{pre}
}

func (nu *NumMatrix) SumRegion(i1 int, j1 int, i2 int, j2 int) int {
	pre := nu.pre
	return pre[i2+1][j2+1] - pre[i1][j2+1] - pre[i2+1][j1] + pre[i1][j1]
}
