func minSubarray(nums []int, p int) int {
	sum:=0
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
	}
	extra:=sum%p
	if extra==0 {
		return extra
	}
	extraM:=make(map[int]int)
	extraM[0]=-1 
	var prev []int
	prev=append(prev,nums[0])
	if prev[0]%p==extra {
		return 1
	}
	rs:=len(nums)
	extraM[prev[0]%p]=0
	for r:=1;r<len(nums);r++{
		if nums[r]%p==extra {
			return 1
		}
		prev=append(prev,prev[r-1]+nums[r])
		l,ok:=extraM[(prev[r]%p-extra+p)%p]
		if ok && r-l<rs {
			rs=r-l
		} 
		extraM[prev[r]%p]=r
	}

	if rs==len(nums) {
		return -1
	}
	return rs
}

