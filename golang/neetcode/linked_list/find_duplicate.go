package linked_list

func findDuplicate1(nums []int) int { // by check loop
	for i, c := range nums {
		if c < 0 {
			continue
		}

		j := i
		for nums[j] >= 0 {
			nums[j], j = -nums[j], nums[j]-1 // mark & jump: index j' == value-1
		}
		if j != i {
			return j + 1 // vj==j+1
		}
	}
	return -1
}

func findDuplicate(nums []int) int { // by store mark into num, num := num*mark(+-1)
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] < 0 {
			return num
		}
		nums[num-1] *= -1
	}
	return -1
}
