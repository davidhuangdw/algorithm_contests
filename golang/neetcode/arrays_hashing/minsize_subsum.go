package arrays_hashing

func minSubArrayLen(target int, nums []int) int {
	sum, j, ans := 0, 0, 0
	for i, v := range nums {
		for j < len(nums) && (j <= i || sum < target) {
			sum += nums[j]
			j++
		}
		if sum >= target && (ans == 0 || j-i < ans) {
			ans = j - i
		}
		sum -= v
	}
	return ans
}
