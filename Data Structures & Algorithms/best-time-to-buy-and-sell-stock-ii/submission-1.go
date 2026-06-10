func maxProfit(prices []int) int {
	lastDay:=len(prices)-1
	profit:=0
	buy:=0;
	sell:=0;
	for i:=1;i<len(prices)-1;i++{
		if prices[i]<prices[i-1] && prices[i]<=prices[i+1] {
			buy=i
		}

		if prices[i]>prices[i-1] && prices[i]>=prices[i+1] {
			sell=i
			profit=profit+prices[sell]-prices[buy]
		}
	}
	if prices[lastDay]>prices[lastDay-1]  {
		sell=lastDay
		profit=profit+prices[sell]-prices[buy]
	}
	return profit
}
