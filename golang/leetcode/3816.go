package main

func lexSmallestAfterDeletion(s string) string {
	cnt, sb := make(map[byte]int), []byte(s)
	for _, b := range sb {
		cnt[b]++
	}
	var ans []byte
	for _, b := range sb {
		for len(ans) > 0 && ans[len(ans)-1] > b && cnt[ans[len(ans)-1]] > 1 {
			cnt[ans[len(ans)-1]]--
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, b)
	}
	for len(ans) > 0 && cnt[ans[len(ans)-1]] > 1 {
		cnt[ans[len(ans)-1]]--
		ans = ans[:len(ans)-1]
	}
	return string(ans)
}
