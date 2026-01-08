package backtrace

import "slices"

func canPartitionKSubsets1(nums []int, k int) bool {
	n, sum := len(nums), 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	sum /= k
	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	vis := make([]bool, n)
	var back func(dep, st, preSum int) bool
	back = func(dep, st, preSum int) bool {
		if dep == k-1 {
			return true
		}
		if preSum == sum {
			return back(dep+1, 0, 0)
		}
		for i := st; i < n; i++ {
			if vis[i] || preSum+nums[i] > sum {
				continue
			}
			vis[i] = true
			if back(dep, i+1, preSum+nums[i]) {
				return true
			}
			vis[i] = false
			if preSum == 0 { // just choose first not-vis for the begin
				break
			}
		}
		return false
	}
	return back(0, 0, 0)
}

func canPartitionKSubsets(nums []int, k int) bool { // by dp -- order unrelated by rule: force append on last group
	n, sum := len(nums), 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	sum /= k
	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	failed := make([]bool, 1<<n)
	var back func(dep, st, preSum, mask int) bool
	// 'mask' contains all other params:
	//    preSum == sum(masked)%seg
	//    dep ==  sum(masked)/seg
	//    st == 1+last-masked-bit
	back = func(dep, st, preSum, mask int) bool {
		if dep == k-1 {
			return true
		}
		if failed[mask] { // dp
			return false
		}
		if preSum == sum {
			return back(dep+1, 0, 0, mask)
		}
		for i := st; i < n; i++ {
			if mask&(1<<i) > 0 || preSum+nums[i] > sum {
				continue
			}
			if back(dep, i+1, preSum+nums[i], mask|(1<<i)) {
				return true
			}
			if preSum == 0 { // just choose first not-vis for the begin
				break
			}
		}
		failed[mask] = true
		return false
	}
	return back(0, 0, 0, 0)
}
