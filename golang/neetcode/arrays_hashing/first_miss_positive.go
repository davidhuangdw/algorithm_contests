package arrays_hashing

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for _, v := range nums {
		for 1 <= v && v <= n && nums[v-1] != v {
			t := nums[v-1]
			nums[v-1], v = v, t
		}
	}
	for v := 1; v <= n; v++ {
		if nums[v-1] != v {
			return v
		}
	}
	return n + 1
}
