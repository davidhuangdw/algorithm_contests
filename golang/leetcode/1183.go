package main

import "slices"

func maximumNumberOfOnes(width int, height int, sideLength int, maxOnes int) int {
	s := sideLength
	freq := make([]int, 0, s*s)
	for i := range s {
		for j := range s {
			freq = append(freq, (height-i+s-1)/s*((width-j+s-1)/s))
		}
	}
	slices.Sort(freq)
	ans := 0
	for _, f := range freq[len(freq)-maxOnes:] {
		ans += f
	}
	return ans
}
