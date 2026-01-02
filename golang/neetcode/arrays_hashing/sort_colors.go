package arrays_hashing

func sortColors1(nums []int) {
	a, l, r := 0, 0, len(nums)-1
	for l <= r {
		switch nums[l] {
		case 0:
			nums[a], nums[l] = nums[l], nums[a]
			a++
			l++
		case 1:
			l++
		case 2:
			nums[l], nums[r] = nums[r], nums[l]
			r--
		}
	}
}

func sortColors2(nums []int) { // partition by sz[]
	var sz [2]int
	for i, _ := range nums {
		for j := 0; j < 2; j++ {
			if nums[i] <= j { // satisfy cond j
				nums[sz[j]], nums[i] = nums[i], nums[sz[j]]
			}
			if nums[sz[j]] <= j { // in case maybe sz[j]==sz[j-1]
				sz[j]++
			}
		}
	}
}

func sortColors(nums []int) { // by count sort
	var cnt [3]int
	for _, v := range nums {
		cnt[v]++
	}
	i := 0
	for v, c := range cnt {
		for ; c > 0; c-- {
			nums[i] = v
			i++
		}
	}
}
