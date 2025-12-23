package stack

func largestRectangleArea(heights []int) int {
	n := len(heights)
	lowerL, lowerR := make([]int, n), make([]int, n)

	for i, h := range heights {
		j := i - 1
		for j >= 0 && heights[j] >= h { // total O(n) since each j skipped only once
			j = lowerL[j]
		}
		lowerL[i] = j
	}
	for i := n - 1; i >= 0; i-- {
		h, j := heights[i], i+1
		for j < n && h <= heights[j] {
			j = lowerR[j]
		}
		lowerR[i] = j
	}

	mx := 0
	for i, h := range heights {
		mx = max(mx, (lowerR[i]-lowerL[i]-1)*h)
	}
	return mx
}
