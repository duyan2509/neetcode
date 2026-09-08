func longestSubarray(nums []int, limit int) int {
	rs:=1
	for r:=1;r<len(nums);r++{
    	l:=r-1
		minNum:=nums[r]
		maxNum:=nums[r]
		for l>=0{
			if nums[l]>maxNum {
				maxNum=nums[l]
			}
			if nums[l]<minNum{
				minNum=nums[l]
			}
			if checkValid(minNum,maxNum,limit) {
				rs=getMax(rs,r-l+1)
			}
			l--
		}
	}
	return rs
}

func checkValid(a int, b int, limit int) bool {
	sum:=a-b
	if sum<=0 {
		sum=-sum
	}
	if sum<=limit {
		return true
	}
	return false
}

func getMax(a int, b int) int {
	if a>b {
		return a
	}
	return b
}