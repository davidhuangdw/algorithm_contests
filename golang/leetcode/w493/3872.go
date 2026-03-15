package main

func longestArithmetic(nums []int) int {
	n := len(nums)
	mx := min(n, 3)

	for i := 0; i < n-1; { // i := [i, i+1] as the first unchange seg for each seq
		j, d := i+2, nums[i+1]-nums[i]
		for ; j < n && nums[j]-nums[j-1] == d; j++ {
		}

		// case 1: ----.----
		k := j + 1 // jump j
		for ; k < n && nums[k]-nums[i] == (k-i)*d; k++ {
		}
		mx = max(mx, min(n, k-i))

		// case 2: ..----
		if i-2 >= 0 && nums[i]-nums[i-2] == 2*d {
			mx = max(mx, min(n, j-i+2))
		}

		i = j - 1 // end of first seq
	}
	return mx
}
