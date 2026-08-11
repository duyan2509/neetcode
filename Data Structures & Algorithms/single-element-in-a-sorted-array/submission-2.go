func singleNonDuplicate(nums []int) int {
	l:=0
	r:=len(nums)-1
	if r==0 || nums[0]!=nums[1]{
		return nums[0]
	}
	if nums[r]!=nums[r-1] {
		return nums[r]
	}
	for l<=r{
		m:=(r+l)/2
		if nums[m]!=nums[m+1] && nums[m]!=nums[m-1] {
			return nums[m]
		} else if (m%2==1 && nums[m]==nums[m-1]) || 
		(m%2==0 && nums[m]==nums[m+1]) {
			l=m+1
		} else if (m%2==1 && nums[m]!=nums[m-1]) || 
		(m%2==0 && nums[m]==nums[m-1]){
			r=m-1
		}
	}
	return nums[l]
}

