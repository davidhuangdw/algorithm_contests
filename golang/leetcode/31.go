package leetcode

func nextPermutation(nums []int)  {
	n := len(nums)
	swap := func(i, j int){
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse := func(l, r int){
		for l < r{
			swap(l, r)
			l += 1
			r -= 1
		}
	}
	i := n-2
	for ; i>=0 && nums[i] >= nums[i+1]; i -=1{}
	if i >= 0{
		for j:=n-1; j>i; j -= 1 {
			if nums[j] > nums[i]{
				swap(i, j)
				break
			}
		}
		reverse(i+1, n-1)
	}else{
		reverse(0, n-1)
	}
}
