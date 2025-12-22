package two_pointers

func trap(height []int) int {
	lmax, rmax, sum := 0, 0, 0
	for l, r := 0, len(height)-1; l <= r; {
		if lmax <= rmax { // c[l] := min(lmax, max(rmax, range[l,r]))==lmax => range [l,r] won't affect l
			lmax = max(lmax, height[l])
			sum += lmax - height[l]
			l++
		} else {
			rmax = max(rmax, height[r])
			sum += rmax - height[r]
			r--
		}
	}
	return sum
}
