package main

func longestBalanced(s string) int {
	c0, n := 0, len(s)
	for i := range n {
		if s[i] == '0' {
			c0++
		}
	}
	mx, ans := 2*min(c0, n-c0), 0
	pos, sum := map[int][]int{0: {-1}}, 0 // bug: forgot zero
	for i := range n {
		if s[i] == '0' {
			sum--
		} else {
			sum++
		}
		for _, d := range []int{-2, 0, 2} {
			tar := sum + d
			for len(pos[tar]) > 0 && i-pos[tar][0] > mx {
				pos[tar] = pos[tar][1:]
			}
			if len(pos[tar]) > 0 {
				ans = max(ans, i-pos[tar][0])
			}
		}
		pos[sum] = append(pos[sum], i)
	}
	return ans
}
