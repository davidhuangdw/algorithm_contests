package arrays_hashing

func majorityElement(nums []int) []int { // return elements more than 1/3
	cnt := make(map[int]int)
	for _, v := range nums {
		if cnt[v] > 0 || len(cnt) < 2 {
			cnt[v]++
			continue
		}
		for x := range cnt {
			cnt[x]--
		}
		for x, c := range cnt {
			if c == 0 {
				delete(cnt, x)
			}
		}
	}
	for x := range cnt {
		cnt[x] = 0
	}
	for _, v := range nums { // 2nd pass: re-count remains
		if _, ok := cnt[v]; ok {
			cnt[v]++
		}
	}

	ans := make([]int, 0)
	for x, c := range cnt {
		if c*3 > len(nums) {
			ans = append(ans, x)
		}
	}
	return ans
}
