func findPeakElement(nums []int) int {
	l:=0
	r:=len(nums)-1
	if r==0 || nums[r]>nums[r-1] {
		return r
	}

	if nums[l]>nums[l+1] {
		return l
	}

	for l<=r {
		m:=(l+r)/2
		if nums[m]>nums[m+1] && nums[m]>nums[m-1]{
			return m
		} else if nums[m]>nums[m-1]{
			l=m+1
		} else {
			r=m
		}
	}
	return l
}
