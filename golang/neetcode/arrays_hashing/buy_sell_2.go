package arrays_hashing

func maxProfit1(prices []int) int {
	buy, sell := int(-1e9), 0
	for _, v := range prices {
		buy, sell = max(buy, sell-v), max(sell, buy+v)
	}
	return sell
}

func maxProfit(prices []int) int { // greedy
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}
