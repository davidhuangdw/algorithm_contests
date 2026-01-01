package dp

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}
	dp := make([]bool, m+1)
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i+j == 0 {
				dp[j] = true
				continue
			}
			c3 := s3[i+j-1]
			dp[j] = (i > 0 && s1[i-1] == c3 && dp[j]) || (j > 0 && s2[j-1] == c3 && dp[j-1])
		}
	}
	return dp[m]
}
