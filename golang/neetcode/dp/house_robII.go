package dp

func robII1(nums []int) int { // nums as cycle
	a0, a1, b0, b1 := 0, 0, 0, 0
	for i, v := range nums {
		switch i {
		case 0:
			a0, a1, b0, b1 = 0, 0, 0, v
		case len(nums) - 1:
			a0, a1, b0, b1 = b0, b1, max(a0+v, b0), max(a0+v, b1)
		default:
			a0, a1, b0, b1 = b0, b1, max(a0+v, b0), max(a1+v, b1)
		}
	}
	return b1
}

func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return max(nums[0], rob(nums[1:]), rob(nums[:len(nums)-1]))
}
