package dp

func change(amount int, coins []int) int {
	cnt := make([]int, amount+1)
	cnt[0] = 1
	for _, c := range coins {
		for i := 0; i+c <= amount; i++ {
			cnt[i+c] += cnt[i]
		}
	}
	return cnt[amount]
}
