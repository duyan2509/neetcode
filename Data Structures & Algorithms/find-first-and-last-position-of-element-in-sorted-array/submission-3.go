func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || (len(nums)==1 && nums[0]!=target) || target>nums[len(nums)-1] || target<nums[0] {
		return []int{-1, -1}
	}
	l:=0
	r:=len(nums)-1
	if nums[0]==target {
		l=0
	} else {
		for l<=r {
			m:=(r+l)/2
			if m>0 && nums[m]==target && nums[m]!=nums[m-1] {
				l=m
				break
			} else if nums[m]>=target {
				r=m-1
			} else {
				l=m+1
			}
		}
	}
	fmt.Printf("l")
	left:=l
	l=0
	if nums[len(nums)-1]==target {
		r=len(nums)-1
	} else {
		r=len(nums)-1
		for l<=r {
			m:=(r+l)/2
			if m<len(nums)-1 && nums[m]==target && nums[m]!=nums[m+1] {
				r=m
				break
			} else if nums[m]>target {
				r=m-1
			} else if nums[m]<=target{
				l=m+1
			}
		}
	}
	if nums[r]!=target {
		r=-1
	}
	if nums[left]!=target {
		left=-1
	}
	return []int{left,r}
}
