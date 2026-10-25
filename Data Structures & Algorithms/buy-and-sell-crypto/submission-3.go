func maxProfit(prices []int) int {
	/*  
		T: O(n)
		S: O(1)
	*/

	max := 0
	
	left := 0
	right := 1
	
	for right < len(prices) {
		buy := prices[left]
		sell := prices[right]
		profit := sell - buy 

		if profit > max {
			max = profit
		} else if sell < buy {
			// found a cheaper buy
			left = right
		}
		right++
	}

	return max
}
