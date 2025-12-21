package two_pointers

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	var ans [][]int
	for i, v := range nums {
		if i > 0 && v == nums[i-1] { // dedup: use lowest i
			continue
		}
		for j, k := i+1, n-1; j < k; j++ {
			if j-1 > i && nums[j] == nums[j-1] { // dedup: lowest j
				continue
			}
			for j < k && nums[j]+nums[k] >= -v {
				k--
			}
			if k+1 < n && nums[j]+nums[k+1] == -v { // dedup: lowest k
				ans = append(ans, []int{v, nums[j], nums[k+1]})
			}
		}
	}
	return ans
}
