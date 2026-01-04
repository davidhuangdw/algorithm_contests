package two_pointers

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = (k%n + n) % n
	for cnt, si := 0, 0; cnt < n; si++ {
		for i := si + k; i != si; i = (i + k) % n {
			nums[i], nums[si] = nums[si], nums[i]
			cnt++
		}
		cnt++
	}
}
