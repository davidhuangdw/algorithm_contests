package sliding_window

func maxSlidingWindow(nums []int, k int) (ans []int) {
	var candQue []int
	ans = make([]int, 0, len(nums)-(k-1))
	if k <= 0 {
		return
	}

	for i, v := range nums {
		for len(candQue) > 0 && nums[candQue[len(candQue)-1]] <= v {
			candQue = candQue[:len(candQue)-1] // pop
		}
		candQue = append(candQue, i)
		if candQue[0] <= i-k {
			candQue = candQue[1:] // pop left
		}
		if i+1-k >= 0 {
			ans = append(ans, nums[candQue[0]])
		}
	}
	return ans
}
