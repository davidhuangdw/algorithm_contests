package greedy

func canJump(nums []int) bool {
	r := 0
	for i, v := range nums {
		if r < i {
			break
		}
		r = max(r, i+v)
	}
	return r >= len(nums)-1

}
