package main

func longestAlternating(nums []int) int {
	ans, n := 1, len(nums)
	up := make([][2]int, n)
	down := make([][2]int, n)
	for i, v := range nums {
		up[i] = [2]int{1, 1}
		down[i] = [2]int{1, 1}
		if i-1 >= 0 && nums[i-1] < v {
			up[i][0] = max(up[i][0], down[i-1][0]+1)
			up[i][1] = max(up[i][0], down[i-1][1]+1)
		}
		if i-1 >= 0 && nums[i-1] > v {
			down[i][0] = max(down[i][0], up[i-1][0]+1)
			down[i][1] = max(down[i][0], up[i-1][1]+1)
		}

		if i-2 >= 0 && nums[i-2] < v {
			up[i][1] = max(up[i][1], down[i-2][0]+1)
		}
		if i-2 >= 0 && nums[i-2] > v {
			down[i][1] = max(down[i][1], up[i-2][0]+1)
		}
		ans = max(ans, up[i][1], down[i][1])
	}
	return ans
}
