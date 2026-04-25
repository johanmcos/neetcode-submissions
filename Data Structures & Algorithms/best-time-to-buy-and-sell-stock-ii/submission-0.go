func maxProfit(prices []int) int {
    profit := 0
    for i := 0; i < len(prices) - 1; i++ {
        profit += max(prices[i+1] - prices[i], 0)
    }

    return profit
}
