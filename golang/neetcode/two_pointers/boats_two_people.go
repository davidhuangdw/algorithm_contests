package two_pointers

import "slices"

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	l, ans := 0, 0
	for r := len(people) - 1; r >= l; r-- {
		ans++
		if l < r && people[l]+people[r] <= limit {
			l++
		}
	}
	return ans
}
