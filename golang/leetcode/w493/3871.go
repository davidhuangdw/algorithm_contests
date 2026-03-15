package main

func countCommas(n int64) (ans int64) {
	for k := int64(1000); k <= n; k *= 1000 {
		ans += n - k + 1
	}
	return ans
}
