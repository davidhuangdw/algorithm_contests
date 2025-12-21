package arrays_hashing

func productExceptSelf(nums []int) []int {
	n := len(nums)
	suf := make([]int, n)
	acc := 1
	for i := n - 1; i >= 0; i-- {
		suf[i] = acc
		acc *= nums[i]
	}

	acc = 1
	for i := 0; i < n; i++ {
		suf[i] = acc * suf[i] // inplace
		acc *= nums[i]
	}
	return suf
}
