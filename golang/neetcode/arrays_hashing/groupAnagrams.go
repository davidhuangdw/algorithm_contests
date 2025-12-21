package arrays_hashing

func groupAnagrams(strs []string) [][]string {
	grp := make(map[[26]int][]string) // dense store by 26-letters
	for _, s := range strs {
		cnt := [26]int{}
		for _, c := range s {
			cnt[c-'a']++
		}
		grp[cnt] = append(grp[cnt], s)
	}
	ans := make([][]string, 0, len(grp))
	for _, o := range grp {
		ans = append(ans, o)
	}
	return ans
}
