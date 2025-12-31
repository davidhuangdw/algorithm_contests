package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	mn, ans := gas[n-1]-cost[n-1], 0
	sum := mn
	for i := 0; i < n-1; i++ {
		d := gas[i] - cost[i]
		sum += d
		if sum < mn {
			ans, mn = i+1, sum
		}
	}
	if sum < 0 {
		return -1
	}
	return ans
}
