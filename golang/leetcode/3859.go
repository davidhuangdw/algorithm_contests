package main

func countSubarrays(nums []int, k int, m int) int64 {
	var ans int64
	n := len(nums)
	var l, r int
	cnt := make(map[int]int)
	less := make(map[int]bool)
	add := func(v, d int) {
		cnt[v] += d
		if cnt[v] <= 0 {
			delete(cnt, v)
			delete(less, v) // bug: also need remove
			return
		}
		if cnt[v] < m {
			less[v] = true
		} else {
			delete(less, v)
		}
	}
	for i := range nums {
		for l < n && len(cnt) < k {
			add(nums[l], 1)
			l++
		}
		for l < n && (len(less) > 0 && cnt[nums[l]] > 0) {
			add(nums[l], 1)
			l++
		}
		r = max(r, l)
		for r < n && cnt[nums[r]] > 0 {
			r++
		}
		if len(less) == 0 && len(cnt) >= k {
			ans += int64(r - l + 1)
		}
		add(nums[i], -1)
	}
	return ans
}
