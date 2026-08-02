type StockSpanner struct {
	prices []int
}

func Constructor() StockSpanner {
	return StockSpanner {
		prices:[]int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	size:=len(this.prices)
	rs:=1
	for i:=size-1;i>=0;i-- {
		if this.prices[i]<=price{
			rs++
		} else {
			break
		}
	}
	this.prices=append(this.prices,price)
	return rs
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 