package arrays_hashing

func subarraySum(nums []int, k int) int {
	pre, cnt, ans := 0, make(map[int]int), 0
	cnt[0] = 1
	for _, v := range nums {
		pre += v
		ans += cnt[pre-k]
		cnt[pre]++
	}
	return ans
}
