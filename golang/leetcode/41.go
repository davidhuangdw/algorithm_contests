package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		for j := nums[i] - 1; 0 <= j && j < n && nums[j] != nums[i]; j = nums[i] - 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	//fmt.Println(nums)
	for v := 1; v <= n; v++ {
		if nums[v-1] != v {
			return v
		}
	}
	return n + 1
}
