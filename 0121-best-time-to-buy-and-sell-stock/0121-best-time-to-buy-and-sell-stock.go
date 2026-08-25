func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// Update minimum buying price
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		// Calculate profit if we sell today
		profit := prices[i] - minPrice

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}