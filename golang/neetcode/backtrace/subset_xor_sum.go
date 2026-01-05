package backtrace

func subsetXORSum1(nums []int) int {
	sum := 0
	for i := 0; i < 1<<len(nums); i++ {
		x, r := 0, i
		for j := 0; r > 0 && j < len(nums); j++ {
			if r&1 > 0 {
				x ^= nums[j]
			}
			r >>= 1
		}
		sum += x
	}
	return sum
}

func subsetXORSum(nums []int) int {
	or := 0
	for _, v := range nums {
		or |= v
	}
	return or << max(0, len(nums)-1) // for each existing bit, half-comb-size is 1, since 1's count == 0's count
}
