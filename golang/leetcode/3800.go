package main

func minimumCost(s string, t string, flipCost int, swapCost int, crossCost int) int64 {
	tol, a := 0, 0
	for i, c := range s {
		if s[i] != t[i] {
			tol++
			if c == '0' {
				a++
			}
		}
	}
	a = min(a, tol-a)
	d := tol - a*2
	u := d / 2 * 2
	return int64(min(a*2*flipCost, swapCost*a) + min(u*flipCost, u/2*(crossCost+swapCost)) + (d-u)*flipCost)
}
