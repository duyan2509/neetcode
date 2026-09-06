func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	currentSum:=0
	currentMax:=0
	l:=0
	rs:=0
	for r:=0;r<len(nums);r++{
		currentSum+=nums[r]
		if nums[r]>currentMax{
			currentMax=nums[r]
		}
		for currentSum+k<(r-l+1)*currentMax {
			currentSum-=nums[l]
			l++
		}
		rs=getMax(rs,r-l+1)
	}	
	return rs
}

func getMax(a int, b int) int {
	if a>b {
		return a
	}
	return b
}
