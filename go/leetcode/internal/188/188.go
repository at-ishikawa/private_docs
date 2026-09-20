package internal

func maxProfit(k int, prices []int) int {
	return maxProfit2(k, prices)
}

// Time complexity: O(n)
// Space complexity: O(k)
func maxProfit2(k int, prices []int) int {
	sold := make([]int, k+1)
	bought := make([]int, k+1)
	for i := 0; i <= k; i++ {
		bought[i] = -1001
	}

	lastDay := len(prices) - 1
	for day := 0; day <= lastDay; day++ {
		price := prices[day]
		for count := k; count >= 1; count-- {
			sold[count] = max(sold[count], bought[count]+price)
			bought[count] = max(bought[count], sold[count-1]-price)
		}
	}
	maxProfit := 0
	for i := 1; i <= k; i++ {
		maxProfit = max(maxProfit, sold[i])
	}
	return maxProfit
}

type stockPrice struct {
	sold   int
	bought int
}

// Time complexity: O(N)
// Space complexity: O(N*k)
func maxProfit1(k int, prices []int) int {
	allDayProfits := make([][]stockPrice, len(prices))
	for day := 0; day < len(prices); day++ {
		allDayProfits[day] = make([]stockPrice, k)
		for transactionCount := 0; transactionCount < k; transactionCount++ {
			allDayProfits[day][transactionCount].sold = -1001
			allDayProfits[day][transactionCount].bought = -1001
		}
	}

	lastDay := len(prices) - 1
	allDayProfits[0][0].bought = -prices[0]
	for day := 1; day <= lastDay; day++ {
		price := prices[day]
		for transactionCount := k - 1; transactionCount >= 0; transactionCount-- {
			allDayProfits[day][transactionCount].sold = max(allDayProfits[day-1][transactionCount].sold, allDayProfits[day-1][transactionCount].bought+price)

			if transactionCount > 0 {
				allDayProfits[day][transactionCount].bought = max(allDayProfits[day-1][transactionCount].bought, allDayProfits[day-1][transactionCount-1].sold-price)
			} else {
				allDayProfits[day][transactionCount].bought = max(allDayProfits[day-1][transactionCount].bought, -price)
			}
		}
	}

	maxProfit := 0
	for i := 0; i < k; i++ {
		maxProfit = max(maxProfit, allDayProfits[lastDay][i].sold)
	}
	return maxProfit
}
