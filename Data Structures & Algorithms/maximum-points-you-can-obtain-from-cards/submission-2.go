func maxScore(cardPoints []int, k int) int {
	rs:=0
	l:=0
	currentSum:=0
	for r:=len(cardPoints)-k;r<len(cardPoints);r++{
		currentSum+=cardPoints[r]
	}
	r:=len(cardPoints)-k
	for l<k{
		if currentSum>rs {
			rs=currentSum
		}
		currentSum=currentSum-cardPoints[r]+cardPoints[l]
		r++
		l++
	}
	if currentSum>rs {
		rs=currentSum
	}
	return rs
}

func calSize(l int, r int, size int) int {
	return size-r+l+1
}