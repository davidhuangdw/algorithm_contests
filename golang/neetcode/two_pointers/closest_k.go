package two_pointers

func findClosestElements1(arr []int, k int, x int) []int {
	n := len(arr)
	if k >= n {
		return arr
	}
	l, r := 0, n-1
	for l <= r {
		md := (l + r) / 2
		if arr[md] <= x {
			l = md + 1
		} else {
			r = md - 1
		}
	}
	l, r = r, r+1
	for r-1-l < k {
		if r >= n || (l >= 0 && x-arr[l] <= arr[r]-x) {
			l--
		} else {
			r++
		}
	}
	return arr[l+1 : r]
}

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) <= k {
		return arr
	}
	l, r := 0, len(arr)-k-1 // lost-nxt, win-nxt
	for l <= r {
		md := (l + r) / 2
		if x-arr[md] > arr[md+k]-x { // lose-nxt: dis(md) > dis(nxt)
			l = md + 1
		} else {
			r = md - 1
		}
	}
	return arr[l : l+k]
}
