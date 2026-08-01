func reverse(x int) int {
	const (
		MaxInt32 = 1<<31 - 1
		MinInt32 = -1 << 31
	)
	sign:=1
	if x<0 {
		sign=-1
		x=-x
	}
	rs:=0
	for x>0 {
		if (10*rs>MaxInt32-(x%10)) || (sign==-1 && -10*rs<MinInt32+(x%10)){
			return 0
		}
		rs=10*rs + (x%10)
		x=x/10
	}
	return rs*sign
}
