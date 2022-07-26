package main


func main() {
	
}

func maxProfit(prices []int) int {
	// prices[i] = prices of stock on i'th day
	// chose a single day to buy one stock
	// choose a different day in the future to buy the stock
	// return max profit
	// track of lowest value
	var lowestValue, highestValue, currentMaxReturn int
	for i := 0 ; i < len(prices) ; i++ {
		if prices[i] > highestValue {
			highestValue = prices[i]
		}
		if prices[i] < lowestValue || i == 0 {
			lowestValue = prices[i]
			// need to reset highest so that it can only be increased afterwards
			highestValue = prices[i]
		}
		if currentMaxReturn < highestValue - lowestValue {
			currentMaxReturn = highestValue - lowestValue
		}
	}
	return currentMaxReturn
}


/*
Runtime: 126 ms, faster than 88.45% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 8 MB, less than 98.76% of Go online submissions for Best Time to Buy and Sell Stock.
Next challenges:
*/
