func search(nums []int, target int) bool {
	l:=0
	r:=len(nums)-1
	m:=0
	for l<=r && r<len(nums){
		if target == nums[l] || target == nums[r] {
			return true
		}
		m=(l+r)/2
		if nums[m]==target {
			return true
		}
		if nums[l]<nums[r] {
			if nums[m]>target {
				r=m-1
			} else {
				l=m+1
			}
		} else if nums[l]==nums[m] && nums[r]==nums[m] {
			l++
		} else {
			if target>nums[l]{
				r=m-1
			}
			if target<nums[l]{
				l=m+1
			}
		}
	}
	return false
}
