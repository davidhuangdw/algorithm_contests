package backtrace

func partition(s string) [][]string {
	ss, ans := []byte(s), make([][]string, 0)
	n := len(ss)
	isPar := make([][]bool, n)
	for i, _ := range isPar {
		isPar[i] = make([]bool, n)
	}
	for l := 1; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if ss[i] == ss[j] && (i+1 >= j-1 || isPar[i+1][j-1]) {
				isPar[i][j] = true
			}
		}
	}

	var dfs func(i int, prev []string)
	dfs = func(i int, prev []string) {
		if i >= n {
			ans = append(ans, append([]string{}, prev...))
			return
		}
		for j := i; j < n; j++ {
			if isPar[i][j] {
				dfs(j+1, append(prev, string(ss[i:j+1])))
			}
		}
	}
	dfs(0, nil)

	//fmt.Printf("%v", ans)
	return ans
}
