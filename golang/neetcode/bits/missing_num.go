package bits

func missingNumber(nums []int) (ans int) {
	for i, v := range nums {
		ans += i + 1 - v
	}
	return
}
