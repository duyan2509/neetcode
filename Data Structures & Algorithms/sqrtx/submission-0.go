func mySqrt(x int) int {
	l:=1
	r:=x
	for l<=r{
		m:=(r+l)/2
		if m*m==x {
			return m
		} else if m*m>x {
			r=m-1
		} else {
			l=m+1
		}
	}
	return l-1
}
