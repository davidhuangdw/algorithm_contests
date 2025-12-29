package dp

func coinChange(coins []int, amount int) int {
	ways := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		ways[i] = -1
	}
	for i := 0; i <= amount; i++ {
		if ways[i] >= 0 {
			for _, c := range coins {
				j := i + c
				if j <= amount && (ways[j] < 0 || ways[i]+1 < ways[j]) {
					ways[j] = ways[i] + 1
				}
			}
		}
	}
	return ways[amount]
}
