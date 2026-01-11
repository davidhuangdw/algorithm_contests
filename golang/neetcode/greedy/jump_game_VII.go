package greedy

func canReach(s string, minJump int, maxJump int) bool {
	// perspective: query each pos by maintain its prev-window that jump to it
	if s[0] != '0' {
		return false
	}
	cnt, n := 0, len(s)
	reach := make([]bool, n)
	reach[0] = true
	for i := 1; i < n; i++ {
		if i-minJump >= 0 && reach[i-minJump] {
			cnt++
		}
		if i-maxJump-1 >= 0 && reach[i-maxJump-1] {
			cnt--
		}
		reach[i] = s[i] == '0' && cnt > 0
	}
	return reach[n-1]
}
