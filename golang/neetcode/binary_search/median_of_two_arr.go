package binary_search

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	less := func(i, j int) bool {
		return j >= m || (i < n && nums1[i] < nums2[j])
	}
	find := func(k int) float64 { // find value of rank k(start from 1)
		if k > n+m {
			return 0
		}
		i, j := 0, 0
		for k > 1 {
			l := k / 2
			r := k - l
			if less(i+l-1, j+r-1) {
				i += l
				k = r
			} else {
				j += r
				k = l
			}
		}
		if less(i, j) {
			return float64(nums1[i])
		}
		return float64(nums2[j])
	}
	if n+m == 0 { // bug
		return 0
	}
	if (n+m)%2 == 0 {
		return (find((n+m)/2) + find((n+m)/2+1)) / 2
	}
	return find((n + m + 1) / 2)
}
