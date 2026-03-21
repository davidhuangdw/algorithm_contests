package main

func numSubmat(mat [][]int) int {
	m, ans := len(mat[0]), 0
	h := make([]int, m)
	for _, row := range mat {
		que := [][3]int{{-1, -1, 0}}
		for j, v := range row {
			if v == 0 {
				h[j] = 0
			} else {
				h[j]++
			}
			for que[len(que)-1][0] >= h[j] {
				que = que[:len(que)-1]
			}
			q := que[len(que)-1]
			i, si := q[1], q[2]
			sj := (j-i)*h[j] + si
			que = append(que, [3]int{h[j], j, sj})
			ans += sj
		}
	}
	return ans
}
