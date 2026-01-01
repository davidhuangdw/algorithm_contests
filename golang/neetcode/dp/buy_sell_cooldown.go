package dp

func maxProfit(prices []int) int {
	INF := int(1e9)
	buy, sellA, sellB := -INF, 0, 0
	for _, v := range prices {
		buy, sellA, sellB = max(buy, sellA-v), sellB, max(sellB, buy+v)
	}
	return sellB
}
