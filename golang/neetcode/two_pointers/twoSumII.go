package two_pointers

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for l, r := 0, n-1; l < r; l++ {
		for l < r && numbers[l]+numbers[r] >= target { // use lowest r
			r--
		}
		if r+1 < n && numbers[l]+numbers[r+1] == target {
			return []int{l + 1, r + 2}
		}
	}
	return nil
}
