package leetcode

func kthDistinct(arr []string, k int) string {
	cnt := map[string]int{}
	for _, s := range arr {
		cnt[s]++
	}
	i := 1
	for _, s := range arr {
		if cnt[s] == 1 {
			if i == k {
				return s
			}
			i ++
		}
	}
	return ""
}