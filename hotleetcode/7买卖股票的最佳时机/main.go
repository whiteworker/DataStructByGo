package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minCost, maxProfit := prices[0], 0
	for _, prices := range prices {
		if prices < minCost {
			minCost = prices
		}
		maxProfit = max(maxProfit, prices-minCost)
	}
	return maxProfit
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	stocks := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(stocks))
}
