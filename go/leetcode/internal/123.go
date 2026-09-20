package internal

import "math"

// Options
// Option 1: Apply recursively
//  1. hasStock: max(maxProfit(prices[i+1], transactionCount, hasStock), prices[i] + maxProfit(prices[i+1], transactionCount-1))
//  2. Otherwise: max(maxProfit(prices[i+1], transactionCount, hasStock), maxProfit(prices[i+1], transactionCount-1, true) - prices[i])
// O(2^n)

// Option 2:
// find 2 combinations of (max_i - min_j) where 0 <= j < i < n
// 2 combinations should be independent, like (j1, i1) < (j2, i2)
// diffs[i][j] = prices[i]-prices[j]

// Option 3:
// Memoization and repeat?
// maxProfits[0][0] = 0
// maxProfits[0][1] = -prices[0]
// maxProfits[1][0] = prices[1]+maxProfits[0][1]
// maxProfits[1][1] = min(maxProfits[0][1], -prices[1])
// maxProfits[2][0] =

// i=0: buy or nothing
// i=1: buy (if nothing on i=0), sell 1st stock, or nothing
// i=2: buy (if sold or nothing on i=1), buy, sell (if buy on i=1), or nothing

func maxProfit(prices []int) int {
	var sold2, sold1 int
	bought1, bought2 := math.MinInt, math.MinInt
	for _, price := range prices {
		sold2 = max(sold2, bought2+price)
		bought2 = max(bought2, sold1-price)
		sold1 = max(sold1, bought1+price)
		bought1 = max(bought1, -price)
	}
	return sold2
}

func maxProfitOriginal(prices []int) int {
	statuses := make([][5]int, len(prices))
	statuses[0] = [5]int{
		0,           // if not bought,
		-prices[0],  // if bought
		math.MinInt, // if sold
		math.MinInt, // if bought again
		math.MinInt, // if sold 2nd stock
	}

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		previous := statuses[i-1]

		array := [5]int{
			previous[0],
			max(previous[1], -price),
			max(
				previous[2],       // sold before
				previous[1]+price, // sold the stock bought before
			),
			math.MinInt,
			math.MinInt,
		}
		// buy 2nd stock
		if i >= 2 {
			array[3] = max(
				previous[3],
				previous[2]-price,
			)
		}
		// sell 2nd stock
		if i >= 3 {
			array[4] = max(
				previous[4],
				previous[3]+price,
			)
		}
		statuses[i] = array
	}

	n := len(prices)
	return max(0, statuses[n-1][2], statuses[n-1][4])
}
