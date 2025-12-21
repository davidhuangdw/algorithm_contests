package arrays_hashing

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	mx := 0
	for _, v := range nums {
		cnt[v]++
		mx = max(mx, cnt[v])
	}

	freq := make(map[int][]int)
	for v, c := range cnt {
		freq[c] = append(freq[c], v)
	}

	ans := make([]int, 0, k)
	for f := mx; f > 0; f-- { // O(n) by iter possible freq-values
		for _, v := range freq[f] {
			ans = append(ans, v)
			if len(ans) == k {
				return ans
			}
		}
	}
	return ans
}
