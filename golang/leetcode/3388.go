package main

func beautifulSplits(nums []int) int {
	n := len(nums)
	lcp := make([][]int, n+1)
	for i := n - 1; i >= 0; i-- {
		lcp[i] = make([]int, n+1)
		for j := n - 1; j > i; j-- {
			if nums[i] == nums[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	ans := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (i <= j-i && lcp[0][i] >= i) || (j-i <= n-j && lcp[i][j] >= j-i) {
				ans++
			}
		}
	}
	return ans
}
