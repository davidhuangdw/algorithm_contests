package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	var que []int
	var ret []int
	for i, v := range nums {
		for len(que) > 0 && que[len(que)-1] < v{
			que = que[:len(que)-1]
		}
		que = append(que, v)
		if i >= k-1 {
			ret = append(ret, que[0])
		}
		if i-k+1 >= 0 && que[0] == nums[i-k+1]{
			que = que[1:]
		}
	}
	return ret
}
