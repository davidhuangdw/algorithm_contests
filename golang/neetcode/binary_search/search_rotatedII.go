package binary_search

func searchR2(nums []int, target int) bool { // allow duplicate
	l, r := 0, len(nums)-1
	for l <= r {
		md := (l + r) / 2
		if nums[md] == target {
			return true
		}

		if nums[md] == nums[l] {
			l++
		} else if (nums[l] < nums[md] && nums[l] <= target && target < nums[md]) ||
			(nums[md] < nums[l] && !(nums[md] < target && target < nums[l])) {
			r = md - 1
		} else {
			l = md + 1
		}
	}
	return false
}
