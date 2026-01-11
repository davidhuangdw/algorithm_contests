package greedy

func maxTurbulenceSize(arr []int) (ans int) {
	a, b := 1, 1
	ans = 1
	for i := 1; i < len(arr); i++ {
		if i%2 == 0 {
			if arr[i] >= arr[i-1] {
				a = 0
			}
			if arr[i] <= arr[i-1] {
				b = 0
			}
		} else {
			if arr[i] <= arr[i-1] {
				a = 0
			}
			if arr[i] >= arr[i-1] {
				b = 0
			}
		}
		a++
		b++
		ans = max(ans, a, b)
	}
	return ans
}
