package binary_search

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		md := (l + r) / 2
		if nums[md] == target {
			return md
		}

		if (target >= nums[0] && nums[0] <= nums[md] && nums[md] < target) ||
			(target < nums[0] && !(target < nums[md] && nums[md] < nums[0])) {
			l = md + 1
		} else {
			r = md - 1
		}

	}
	return -1
}
